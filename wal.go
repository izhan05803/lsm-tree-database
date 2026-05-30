package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// WALEntry represents a single operation in the Write-Ahead Log
type WALEntry struct {
	Op    string // "PUT" or "DELETE"
	Key   string
	Value string
}

// WAL represents the Write-Ahead Log for durability
type WAL struct {
	file   *os.File
	writer *bufio.Writer
	mu     sync.Mutex
	path   string
}

// NewWAL creates or opens a WAL file
func NewWAL(path string) (*WAL, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return &WAL{
		file:   file,
		writer: bufio.NewWriter(file),
		path:   path,
	}, nil
}

// Append writes a new entry to the WAL
// Critical: must flush to disk immediately for durability
func (w *WAL) Append(op string, key string, value string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := WALEntry{
		Op:    op,
		Key:   key,
		Value: value,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = w.writer.Write(data)
	if err != nil {
		return err
	}

	// Write newline to separate entries
	err = w.writer.WriteByte('\n')
	if err != nil {
		return err
	}

	// Flush to disk to ensure durability
	return w.writer.Flush()
}

// Read reads all entries from the WAL file
func (w *WAL) Read() ([]WALEntry, error) {
	file, err := os.Open(w.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []WALEntry{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var entries []WALEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var entry WALEntry
		err := json.Unmarshal([]byte(line), &entry)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, scanner.Err()
}

// Close closes the WAL file
func (w *WAL) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		return w.file.Close()
	}
	return nil
}
