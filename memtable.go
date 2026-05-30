package main

import (
	"fmt"
	"sort"
)

// Pair represents a key-value pair
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// MemTable is an in-memory store using a hash map
// It provides O(1) Put, Get, Delete operations
type MemTable[K comparable, V any] struct {
	data map[K]V // The actual storage
	size int     // Number of entries (for flush threshold)
}

// NewMemTable creates a new empty MemTable
func NewMemTable[K comparable, V any]() *MemTable[K, V] {
	return &MemTable[K, V]{
		data: make(map[K]V),
		size: 0,
	}
}

// Put adds or updates a key-value pair
// If key is NEW: increment size
// If key EXISTS: just update value, size stays same
func (m *MemTable[K, V]) Put(key K, value V) {
	_, exists := m.data[key]
	m.data[key] = value
	if !exists {
		m.size++
	}
}

// Get retrieves a value by key
// Returns (value, true) if found
// Returns (zero-value, false) if not found
func (m *MemTable[K, V]) Get(key K) (V, bool) {
	value, ok := m.data[key]
	if !ok {
		var zero V
		return zero, false
	}
	return value, true
}

// Delete removes a key-value pair
// Only decrements size if key actually existed
func (m *MemTable[K, V]) Delete(key K) {
	if _, ok := m.data[key]; ok {
		delete(m.data, key)
		m.size--
	}
}

// Size returns the current number of entries
func (m *MemTable[K, V]) Size() int {
	return m.size
}

// Flush returns all entries sorted by key and clears the MemTable
// This is called when flushing to disk (creating an SSTable)
func (m *MemTable[K, V]) Flush() []Pair[K, V] {
	// Convert map to slice of Pairs
	pairs := make([]Pair[K, V], 0, len(m.data))
	for k, v := range m.data {
		pairs = append(pairs, Pair[K, V]{Key: k, Value: v})
	}

	// Sort by key (using fmt.Sprint for generic comparison)
	sort.Slice(pairs, func(i, j int) bool {
		return fmt.Sprint(pairs[i].Key) < fmt.Sprint(pairs[j].Key)
	})

	// Reset the MemTable (frees memory, prepares for reuse)
	m.data = make(map[K]V)
	m.size = 0

	return pairs
}
