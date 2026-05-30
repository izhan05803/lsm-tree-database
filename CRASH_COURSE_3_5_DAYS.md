# LSM Tree Database - 3-5 Day Crash Course

Build a working LSM database + interview prep in 3-5 days. No fluff, just essentials.

---

## DAY 1: Understanding & Setup (6-8 hours)

### Morning (2-3 hours): Theory
**What is LSM Tree?**
- Log-Structured Merge Tree
- Optimized for fast writes (append-only operations)
- Trade-off: Slower reads (multiple files to check)
- Why? Real-world use: RocksDB, LevelDB, Cassandra

**Why not B-Tree?**
- B-Tree: Random I/O per write = slow
- LSM: Append I/O per write = fast (10-100x faster)

**Core Idea:**
```
Write Path: Client → WAL (disk) → MemTable (memory)
           → When full → SSTable (disk)
           
Read Path: Client → MemTable → SSTable1 → SSTable2 → ...

Background: Many SSTables → Compaction → Fewer SSTables
```

**The 5 Components:**
1. **MemTable** - Fast in-memory storage (Go map)
2. **WAL** - Write-ahead log (durability)
3. **SSTable** - Immutable disk files (sorted data)
4. **Compaction** - Merge old files
5. **Database** - Orchestrates everything

### Afternoon (3-5 hours): Design & Setup

**Draw on paper:**
- Write path diagram
- Read path diagram
- Compaction process

**Design decisions:**
- MemTable flush size: 64MB
- Compaction trigger: 10 SSTables
- File format: Simple binary (key_len, key, val_len, value)

**Setup Go project:**
```bash
cd D:\embedded-lsm-tree-engine
go mod init lsm-db
```

**End of Day 1:** You understand LSM, have basic design, project setup done.

---

## DAY 2: Core Components (8-10 hours)

### Morning (4 hours): MemTable + WAL

**Implement MemTable (memtable.go):**
```go
type MemTable struct {
    data map[string]string
    size int
}

func (m *MemTable) Put(key, value string) {}
func (m *MemTable) Get(key string) (string, bool) {}
func (m *MemTable) Delete(key string) {}
func (m *MemTable) Size() int { return m.size }
func (m *MemTable) Entries() []Pair {} // sorted
```

**Implement WAL (wal.go):**
```go
type WALEntry struct {
    Key   string
    Value string
    Op    string // "put" or "delete"
}

func (w *WAL) Append(entry WALEntry) error {}
func (w *WAL) Read() ([]WALEntry, error) {}
```

**Test both locally** - Print output to verify they work.

### Afternoon (4-6 hours): SSTable

**Implement SSTable (sstable.go):**
```go
type SSTable struct {
    entries []Pair // sorted key-value pairs
    path    string
}

func (s *SSTable) Get(key string) (string, error) {}
func (s *SSTable) Write() error {} // serialize to file
func (s *SSTable) Read(path string) error {} // deserialize from file
```

**Simple binary format:**
```
[num_entries:4 bytes][key1_len][key1][val1_len][val1][key2_len][key2]...
```

**End of Day 2:** MemTable, WAL, SSTable all working independently.

---

## DAY 3: Database Core (8-10 hours)

### Morning (4 hours): Database Structure

**Implement Database (db.go):**
```go
type DB struct {
    memtable  *MemTable
    wal       *WAL
    sstables  []*SSTable
    path      string
    maxSize   int // flush threshold
    threshold int // compaction threshold
}

func NewDB(path string) (*DB, error) {}
func (db *DB) Put(key, value string) error {}
func (db *DB) Get(key string) (string, error) {}
func (db *DB) Close() error {}
```

**Write Path:**
```go
func (db *DB) Put(key, value string) error {
    // 1. Write to WAL
    db.wal.Append(WALEntry{Key: key, Value: value, Op: "put"})
    
    // 2. Write to MemTable
    db.memtable.Put(key, value)
    
    // 3. Check if flush needed
    if db.memtable.Size() > db.maxSize {
        db.flush()
    }
    return nil
}
```

**Read Path:**
```go
func (db *DB) Get(key string) (string, error) {
    // Check memtable first
    if val, ok := db.memtable.Get(key); ok {
        return val, nil
    }
    
    // Check SSTables (newest first)
    for i := len(db.sstables) - 1; i >= 0; i-- {
        val, err := db.sstables[i].Get(key)
        if err == nil {
            return val, nil
        }
    }
    return "", NotFound
}
```

### Afternoon (4-6 hours): Compaction + Finalize

**Implement Compaction (compaction.go):**
```go
func (db *DB) Compact() error {
    if len(db.sstables) < db.threshold {
        return nil // nothing to compact
    }
    
    // Merge first 2 SSTables
    merged := merge(db.sstables[0], db.sstables[1])
    
    // Write merged SSTable
    merged.Write()
    
    // Remove old ones, add merged
    db.sstables = append(db.sstables[2:], merged)
    
    return nil
}
```

**Test everything together:**
- Put 100 values
- Get them back
- Verify persistence

**End of Day 3:** Full database working (Put/Get/Delete/Compaction).

---

## DAY 4: Polish & Testing (6-8 hours)

### Morning (3-4 hours): Testing & Verification

**Write quick tests:**
```go
func TestPutGet(t *testing.T) {
    db := NewDB("test_db")
    db.Put("key1", "value1")
    val, _ := db.Get("key1")
    assert(val == "value1")
}

func TestCompaction(t *testing.T) {
    // Add many values, trigger compaction
    // Verify all values still readable
}

func TestRecovery(t *testing.T) {
    // Close DB, reopen, verify data persists
}
```

**Stress test:**
```go
// Add 10,000 values
for i := 0; i < 10000; i++ {
    db.Put(fmt.Sprintf("key%d", i), "value")
}

// Read them all back
for i := 0; i < 10000; i++ {
    val, _ := db.Get(fmt.Sprintf("key%d", i))
    assert(val == "value")
}
```

### Afternoon (3-4 hours): Documentation & Interview Prep

**Write README.md:**
```markdown
# LSM Tree Database

A simple LSM Tree implementation in Go.

## Architecture
- MemTable: In-memory storage
- WAL: Durability
- SSTable: Disk storage
- Compaction: Background maintenance

## API
- Put(key, value)
- Get(key)
- Delete(key)
- Close()

## Performance
- Write: ~100K ops/sec
- Read: ~50K ops/sec
```

**Prepare interview pitch (60 seconds):**
> "I built an LSM Tree database to understand write-optimized storage. LSM Trees achieve fast writes by appending to memory and log instead of random disk I/O. When the in-memory table fills, I flush it to a disk file (SSTable). Reads are slower because they check multiple files, but background compaction merges files to improve read performance. The key insight is batching writes in memory for speed, with a durability guarantee through the write-ahead log."

**End of Day 4:** Working database + basic documentation + interview pitch.

---

## DAY 5 (Optional): Interview Mastery (4-6 hours)

### Morning (2-3 hours): Answer Key Questions

**Q1: "What is an LSM Tree?"**
A: Log-Structured Merge Tree. It optimizes writes by appending data (fast) rather than random I/O (slow). Reads are slower because we check multiple sorted files. Background compaction merges files.

**Q2: "Why LSM instead of B-Tree?"**
A: B-Trees do random I/O per write. LSM batches writes in memory then appends sequentially. ~100x faster writes. Trade-off: slower reads.

**Q3: "Walk me through a write operation."**
A: 
1. Write to WAL (append-only, durable to disk)
2. Insert in MemTable (fast, in-memory)
3. If MemTable > size_threshold, flush to SSTable
4. Return success

**Q4: "Walk me through a read operation."**
A:
1. Check MemTable (O(1) hash lookup) - FASTEST
2. If not found, check SSTables newest-first
3. Each SSTable: binary search (O(log n))
4. Return value or not found

**Q5: "What happens if we crash?"**
A: WAL is durable to disk. On restart, we replay WAL to recover lost writes. No data loss because every write was logged before going to memory.

**Q6: "What's the read amplification?"**
A: Number of files we might check. If 10 SSTables, worst case is O(10 log n). Solution: bloom filters for fast negative checks, caching for hot data.

**Q7: "How does compaction work?"**
A: When SSTable count exceeds threshold, merge oldest ones. Reads all data from multiple files, writes to single new file, deletes old files. Reduces files and improves read speed.

**Q8: "Why is SSTable immutable?"**
A: Immutability simplifies concurrency (no locking), enables efficient merging during compaction, and allows background operations without blocking reads.

### Afternoon (2-3 hours): Code Walkthrough Practice

**Practice explaining:**
1. MemTable implementation (2 min)
2. WAL append/recover (2 min)
3. Database Put operation (3 min)
4. Compaction merge (2 min)

**Record yourself once.** Did you explain it clearly? Did you hesitate? Practice again.

**End of Day 5:** Fully prepared for interview.

---

## Interview Checklist

- [ ] Can draw LSM architecture in 5 minutes
- [ ] Can explain why LSM is write-optimized
- [ ] Can walk through Put/Get/Compaction
- [ ] Have answers to 8 key questions ready
- [ ] Code compiles and all tests pass
- [ ] Can explain every line of your code
- [ ] Have 1-2 debugging stories ready
- [ ] Practiced 60-second pitch 5+ times
- [ ] Feel confident in your understanding

---

## Quick Reference: Code Structure

```
db.go
├── Put(key, value) 
│   ├── wal.Append()
│   ├── memtable.Put()
│   └── flush() if needed
├── Get(key)
│   ├── memtable.Get()
│   └── sstable.Get() for each
└── Close()

memtable.go
├── Put(key, value)
├── Get(key)
└── Entries() // sorted

wal.go
├── Append(entry)
└── Read() // for recovery

sstable.go
├── Write() // serialize
└── Read() // deserialize
└── Get(key) // binary search

compaction.go
├── merge(sstable1, sstable2)
└── removeDeadSpace()
```

---

## Performance Goals

- **Write:** 1000+ ops (before compaction triggers)
- **Read (MemTable hit):** <1ms
- **Read (SSTable hit):** <10ms
- **Compaction:** Should not block writes

---

## Minimum Code Size

Don't overthink. This should be ~500 lines total:
- MemTable: ~50 lines
- WAL: ~80 lines
- SSTable: ~100 lines
- Database: ~150 lines
- Compaction: ~50 lines
- Tests: ~70 lines

If you have more, simplify.

---

## If You Only Have 2 Days

Skip Day 4 & 5. Focus on:
- Day 1: Theory + Design
- Day 2: Implement MemTable + WAL + SSTable
- Day 3: Database + Compaction + Basic tests

Then memorize the Quick Reference section for interview.

---

## If You Have 1 Day (Panic Mode)

1. Understand LSM concept (30 min)
2. Implement MemTable only (1 hour)
3. Implement Database wrapper (1 hour)
4. Write basic tests (30 min)
5. Memorize the 8 key Q&A (1 hour)
6. Practice pitch (30 min)

Not ideal, but better than nothing.

---

## Most Important

Your goal isn't perfect code. Your goal is:
1. Working database (basic Put/Get)
2. Understanding why LSM works this way
3. Ability to explain it clearly in interview

Focus on #2 and #3. #1 will follow naturally.

Good luck. You've got this.
