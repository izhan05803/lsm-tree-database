package main

import (
	"encoding/gob"
	"errors"
	"io"
	"os"
)

// Constants for SSTable
const TOMBSTONE = "<<DELETED>>"

// Errors
var (
	ErrNotFound = errors.New("key not found")
	ErrDeleted  = errors.New("key was deleted")
)

// SSTable represents a Sorted String Table on disk
type SSTable[K comparable, V any] struct {
	path  string
	pairs []Pair[K, V]
}

// NewSSTable creates a new SSTable from a file
func NewSSTable[K comparable, V any](path string) (*SSTable[K, V], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var pairs []Pair[K, V]
	for {
		var pair Pair[K, V]
		if err := decoder.Decode(&pair); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		pairs = append(pairs, pair)
	}

	return &SSTable[K, V]{
		path:  path,
		pairs: pairs,
	}, nil
}

// writeSSTable converts a MemTable to an SSTable file
func writeSSTable[K comparable, V any](memtable *MemTable[K, V], path string) (*SSTable[K, V], error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pairs := memtable.Flush()

	encoder := gob.NewEncoder(file)
	for _, pair := range pairs {
		if err := encoder.Encode(pair); err != nil {
			return nil, err
		}
	}

	return &SSTable[K, V]{path: path, pairs: pairs}, nil
}

// Get retrieves a value from SSTable by key
func (s *SSTable[K, V]) Get(key K) (V, error) {
	file, err := os.Open(s.path)
	if err != nil {
		var zero V
		return zero, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	for {
		var pair Pair[K, V]
		if err := decoder.Decode(&pair); err != nil {
			if err == io.EOF {
				break
			}
			var zero V
			return zero, err
		}

		keyInDB := any(pair.Key).(string)
		if keyInDB == any(key).(string) {
			if any(pair.Value).(string) == TOMBSTONE {
				var zero V
				return zero, ErrDeleted
			}
			return pair.Value, nil
		}

		if keyInDB > any(key).(string) {
			var zero V
			return zero, ErrNotFound
		}
	}

	var zero V
	return zero, ErrNotFound
}
