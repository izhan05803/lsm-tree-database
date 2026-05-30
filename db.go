package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	FLUSH_THRESHOLD   = 20
	COMPACT_THRESHOLD = 5
	DB_DIR            = "./data"
	WAL_FILE          = "./data/wal.log"
)

type DB[K comparable, V any] struct {
	memtable  *MemTable[K, V]
	wal       *WAL
	sstables  []*SSTable[K, V]
	sstableID int
}

func NewDB[K comparable, V any]() (*DB[K, V], error) {
	err := os.MkdirAll(DB_DIR, 0755)
	if err != nil {
		return nil, err
	}

	wal, err := NewWAL(WAL_FILE)
	if err != nil {
		return nil, err
	}

	db := &DB[K, V]{
		memtable:  NewMemTable[K, V](),
		wal:       wal,
		sstables:  make([]*SSTable[K, V], 0),
		sstableID: 0,
	}

	err = db.recover()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// recover replays WAL entries to restore state
func (db *DB[K, V]) recover() error {
	entries, err := db.wal.Read()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		switch entry.Op {
		case "PUT":
			key := any(entry.Key).(K)
			value := any(entry.Value).(V)
			db.memtable.Put(key, value)
		case "DELETE":
			key := any(entry.Key).(K)
			db.memtable.Delete(key)
		}
	}

	return nil
}

func (db *DB[K, V]) Put(key K, value V) error {
	err := db.wal.Append("PUT", fmt.Sprint(key), fmt.Sprint(value))
	if err != nil {
		return err
	}

	db.memtable.Put(key, value)

	if db.memtable.Size() > FLUSH_THRESHOLD {
		err := db.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB[K, V]) Get(key K) (V, error) {
	if val, ok := db.memtable.Get(key); ok {
		return val, nil
	}

	for i := len(db.sstables) - 1; i >= 0; i-- {
		val, err := db.sstables[i].Get(key)
		if err == ErrDeleted {
			var zero V
			return zero, errors.New("key was deleted")
		}
		if err == nil {
			return val, nil
		}
	}

	var zero V
	return zero, errors.New("key not found")
}

func (db *DB[K, V]) Delete(key K) error {
	err := db.wal.Append("DELETE", fmt.Sprint(key), "")
	if err != nil {
		return err
	}

	db.memtable.Delete(key)
	return nil
}

// Flush converts MemTable to SSTable on disk
func (db *DB[K, V]) Flush() error {
	if db.memtable.Size() == 0 {
		return nil
	}

	sstablePath := filepath.Join(DB_DIR, fmt.Sprintf("sstable_%d.sst", db.sstableID))
	db.sstableID++

	sstable, err := writeSSTable(db.memtable, sstablePath)
	if err != nil {
		return err
	}

	db.sstables = append(db.sstables, sstable)

	if len(db.sstables) > COMPACT_THRESHOLD {
		err := db.Compact()
		if err != nil {
			return err
		}
	}

	return nil
}

// Compact merges the first 2 SSTables
func (db *DB[K, V]) Compact() error {
	if len(db.sstables) < 2 {
		return nil
	}

	table1 := db.sstables[0]
	table2 := db.sstables[1]

	pairs1 := table1.pairs
	pairs2 := table2.pairs

	allPairs := append(pairs1, pairs2...)

	mergedMemTable := NewMemTable[K, V]()
	for _, pair := range allPairs {
		mergedMemTable.Put(pair.Key, pair.Value)
	}

	mergedPath := filepath.Join(DB_DIR, fmt.Sprintf("sstable_%d.sst", db.sstableID))
	db.sstableID++

	mergedSSTable, err := writeSSTable(mergedMemTable, mergedPath)
	if err != nil {
		return err
	}

	os.Remove(table1.path)
	os.Remove(table2.path)

	newSSTables := make([]*SSTable[K, V], 0)
	for i := 2; i < len(db.sstables); i++ {
		newSSTables = append(newSSTables, db.sstables[i])
	}
	newSSTables = append(newSSTables, mergedSSTable)

	db.sstables = newSSTables

	return nil
}

// Close closes the database
func (db *DB[K, V]) Close() error {
	return db.wal.Close()
}
