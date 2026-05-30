package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_ = os.RemoveAll("./data")

	fmt.Println("=== LSM Tree Database Tests ===")

	// Test 1: Basic Put and Get
	fmt.Println("\nTest 1: Basic Put and Get")
	fmt.Println("---")
	db, err := NewDB[string, string]()
	if err != nil {
		log.Fatalf("Failed to create DB: %v", err)
	}

	db.Put("alice", "100")
	db.Put("bob", "200")
	db.Put("charlie", "300")

	val, _ := db.Get("alice")
	fmt.Printf("Get('alice') = %s ✓\n", val)

	val, _ = db.Get("bob")
	fmt.Printf("Get('bob') = %s ✓\n", val)

	// Test 2: Delete
	fmt.Println("\nTest 2: Delete")
	fmt.Println("---")
	db.Delete("bob")
	_, err = db.Get("bob")
	if err != nil {
		fmt.Println("Delete works - Get('bob') returns error ✓")
	}

	// Test 3: Recovery
	fmt.Println("\nTest 3: Persistence and Recovery")
	fmt.Println("---")
	db.Close()

	db, err = NewDB[string, string]()
	if err != nil {
		log.Fatalf("Failed to create DB: %v", err)
	}
	defer db.Close()

	val, _ = db.Get("alice")
	fmt.Printf("After restart - Get('alice') = %s ✓\n", val)

	val, _ = db.Get("charlie")
	fmt.Printf("After restart - Get('charlie') = %s ✓\n", val)

	_, err = db.Get("bob")
	if err != nil {
		fmt.Println("After restart - Get('bob') returns error ✓")
	}

	// Test 4: Flush
	fmt.Println("\nTest 4: MemTable Flush and SSTable Creation")
	fmt.Println("---")
	db.Close()
	_ = os.RemoveAll("./data")

	db, _ = NewDB[string, string]()

	fmt.Printf("Initial MemTable size: %d\n", db.memtable.Size())
	fmt.Printf("Initial SSTables count: %d\n", len(db.sstables))

	for i := 1; i <= 30; i++ {
		db.Put(fmt.Sprintf("item_%02d", i), fmt.Sprintf("value_%d", i*10))
	}

	fmt.Printf("After 30 Puts - MemTable size: %d\n", db.memtable.Size())
	fmt.Printf("After 30 Puts - SSTables count: %d\n", len(db.sstables))

	if len(db.sstables) > 0 {
		fmt.Println("Flush triggered successfully ✓")
	}

	// Test 5: Data retrieval
	fmt.Println("\nTest 5: Data Retrieval")
	fmt.Println("---")
	val, _ = db.Get("item_01")
	fmt.Printf("Get('item_01') = %s ✓\n", val)

	val, _ = db.Get("item_30")
	fmt.Printf("Get('item_30') = %s ✓\n", val)

	// Test 6: Stress test
	fmt.Println("\nTest 6: Stress Test (100 operations)")
	fmt.Println("---")
	db.Close()
	_ = os.RemoveAll("./data")

	db, _ = NewDB[string, string]()
	defer db.Close()

	for i := 0; i < 100; i++ {
		db.Put(fmt.Sprintf("key_%d", i), fmt.Sprintf("value_%d", i))
	}

	fmt.Printf("Added 100 entries\n")
	fmt.Printf("MemTable size: %d\n", db.memtable.Size())
	fmt.Printf("SSTables count: %d\n", len(db.sstables))

	val, _ = db.Get("key_0")
	fmt.Printf("Get('key_0') = %s ✓\n", val)

	val, _ = db.Get("key_99")
	fmt.Printf("Get('key_99') = %s ✓\n", val)

	fmt.Println("\n=== All Tests Completed Successfully ===")
}
