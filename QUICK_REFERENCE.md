# Quick Reference: Interview Talking Points

## 60-Second Overview

> "I built an LSM Tree database from scratch to understand how modern databases achieve high write throughput. LSM Trees optimize for writes by batching them in memory (memtable), then writing to disk sequentially. This is much faster than random disk I/O in B-Trees. The trade-off is that reads are slower because we need to check multiple files. I implemented all core components: in-memory memtable for fast writes, write-ahead log for durability, sorted string tables for disk storage, and background compaction to manage file growth. The project taught me about systems design trade-offs, concurrency patterns, and failure recovery."

---

## 5-Minute Deep Dive

### Architecture Overview (1 min)
- **MemTable**: In-memory red-black tree or skip list. When it reaches 64MB (configurable), we flush it to disk.
- **Write-Ahead Log**: Append-only log file. Every write is logged before going to memtable. If we crash, we replay the WAL during recovery.
- **SSTable**: Immutable sorted key-value file on disk. Binary format with index at the end for O(log n) binary search.
- **Compaction**: When SSTable count exceeds 10 (configurable), we merge the oldest ones. This reduces files and improves read performance.
- **Manifest**: JSON file tracking which SSTables are valid and database version.

### Write Path (1 min)
```
Put(key, value) →
  1. Append to WAL (durable to disk)
  2. Insert into MemTable (fast, in-memory)
  3. Check if size > maxMemtableSize
  4. If yes: Flush MemTable to SSTable, start new MemTable
  5. Return success
```

### Read Path (1 min)
```
Get(key) →
  1. Check MemTable first (O(1) map lookup) ← FASTEST
  2. If not found, check SSTables in order (newest first)
  3. Each SSTable: Binary search in index (O(log n))
  4. Return value or not found
```

### Compaction (1 min)
```
Background process:
  Check if #SSTables > compactionThreshold (10)
  If yes:
    Select oldest SSTables (e.g., first 2)
    Merge them (read both, write merged output)
    Remove duplicate keys, delete tombstones
    Delete old SSTables, update manifest
  Result: Fewer files, faster reads, reclaimed space
```

### Recovery (1 min)
```
On startup:
  1. Read manifest (which SSTables are valid)
  2. Replay WAL (reconstructs memtable to crash point)
  3. No data loss because WAL logged everything
  4. Discard incomplete operations from crashed memtable
```

---

## Common Questions & Answers

### "Why LSM Tree instead of B-Tree?"

**B-Tree writes:** Random I/O (slow)
- Balanced tree structure requires shuffling data at multiple levels
- Each write might require 3-5 disk seeks
- ~1000-5000 writes/sec

**LSM Tree writes:** Sequential I/O (fast)
- All writes are appends to memory + log (sequential)
- No tree rebalancing on every write
- ~100,000+ writes/sec (10-100x faster)

**Trade-off:** Slower reads (check multiple files) vs much faster writes

**When to use:** Write-heavy workloads (logs, events, time-series data)

---

### "How is an LSM Tree write-optimized?"

Three key reasons:

1. **Sequential I/O**: Append to WAL (sequential) beats random seeks
2. **Batching**: MemTable batches writes in memory (millions/sec)
3. **Deferred Updates**: Updates don't touch existing disk files; they're merged during compaction

Real-world: RocksDB, LevelDB, Cassandra all use LSM for this reason.

---

### "What's the read amplification?"

Read amplification = number of files you might need to check

**Best case:** Data in MemTable (1 lookup)
**Worst case:** Data in oldest SSTable (need to check 10+ files)

**Solution:** 
- Bloom filters (fast negative checks)
- Caching (cache hot SSTable blocks)
- Tiered compaction (limit levels to ~3-5)

---

### "What happens if you crash during compaction?"

**Scenario**: Merging SSTables 1+2 into 3, then server crashes

**Solution:**
1. Manifest tracks which SSTables are "current"
2. During compaction, don't update manifest until merge completes
3. On crash, manifest still points to old SSTables 1+2
4. Partial file 3 is garbage collected
5. Data is safe; just restart

---

### "How do you handle deleted keys?"

**Tombstone approach** (used in RocksDB, Cassandra):
1. Delete operation writes special "tombstone" marker
2. On read, if you see tombstone, return not found
3. During compaction, if you see tombstone, actually delete the entry
4. Eventually (after compaction), old deleted data disappears

**Alternative**: Mark with timestamp, gc after TTL

---

### "What if disk is full?"

**Handling:**
1. Try to write returns error
2. Application should see "disk full" error
3. Could trigger immediate compaction to free space
4. Or reject writes with proper error message
5. Operator needs to add more disk or increase retention

---

### "How does concurrency work?"

**Single-threaded approach** (what I did first):
- Simple, no race conditions
- Good for learning

**Production approach** (RocksDB):
- RWMutex on memtable (many readers, exclusive writer)
- Background goroutine for compaction (doesn't block writes)
- SSTable access is read-only (inherently safe)
- Careful CAS operations for version management

---

### "Memtable overflow handling?"

When MemTable reaches max size:
1. Create new "immutable memtable"
2. Start writing to fresh "active memtable"
3. Background thread flushes immutable memtable to SSTable
4. This allows reads to hit immutable memtable while we flush

**Without this:** Write stalls during flush (bad UX)

---

## Technical Details to Have Ready

### Performance Characteristics

```
Operation      MemTable Hit    SSTable Hit      Worst Case
Put            O(1)            N/A              O(1)
Get            O(1)            O(k log n)       O(k log n)
Delete         O(1)            N/A              O(1)
Scan           O(m)            O(k log n + m)   O(k log n + m)

k = number of SSTables (~10)
n = entries per SSTable (~1M)
m = results returned
```

### File Sizes in Practice
- MemTable flush threshold: 64MB
- SSTable file size: 64MB-512MB
- WAL file: 64MB or less
- Manifest: <1MB (just metadata)

### Compaction Strategies
- **Leveled**: RocksDB style, bounded LSM levels, best reads
- **Tiered**: Simpler, uses more space, still OK reads
- I implemented: Tiered (easier to understand)

---

## Red Flags to Avoid

**Don't say:**
- "I just copied RocksDB code" ❌ (Show understanding, not copying)
- "I don't know why it works" ❌ (You should understand your code)
- "I didn't think about crashes" ❌ (Recovery is critical)
- "I never tested concurrent access" ❌ (Real systems are concurrent)

**Do say:**
- "I made a trade-off here: X vs Y" ✅
- "I tested this scenario..." ✅
- "I identified this bottleneck and fixed it by..." ✅
- "In production, I would add..." ✅

---

## Impressive Follow-Up Answers

### "What would you do differently?"

**If more time:**
1. Add bloom filters for faster negative lookups (reduce disk I/O 10-100x)
2. Implement compression (save 50-70% disk space)
3. Add block caching (reduce read latency)
4. Implement leveled compaction (better read amplification)
5. Add metrics/observability (understand performance in production)
6. Concurrent multi-threaded writes with MVCC
7. Replication and distributed consistency

### "What was the hardest part?"

**Good answer:**
"Testing crash recovery. I had to think through all failure modes:
- Crash during memtable flush
- Crash during compaction
- Crash with partial WAL write
- Corrupt manifest file

I solved this by writing a test harness that simulates crashes at specific points, then verifies data integrity. This gave me confidence that recovery actually works."

### "How would you scale this?"

**Distributed version:**
1. Partition data by key range (sharding)
2. Each shard is an LSM database on a different server
3. Write to primary, replicate to secondaries
4. Use consensus (Raft) for replica coordination
5. Client routes based on key partition

---

## Code Snippets to Memorize

### Put Operation (Core Logic)
```go
func (db *DB) Put(key string, value string) error {
    // 1. Write to WAL for durability
    db.wal.Append(key, value)
    
    // 2. Write to memtable
    db.memtable.Put(key, value)
    
    // 3. Check if flush needed
    if db.memtable.Size() > db.maxMemtableSize {
        db.flush()
    }
    return nil
}
```

### Get Operation (Core Logic)
```go
func (db *DB) Get(key string) (string, error) {
    // 1. Check memtable (fast)
    if val, exists := db.memtable.Get(key); exists {
        return val, nil
    }
    
    // 2. Check SSTables (newest first)
    for i := len(db.sstables) - 1; i >= 0; i-- {
        if val, exists := db.sstables[i].Get(key); exists {
            return val, nil
        }
    }
    
    return "", NotFound
}
```

### Compaction Merge (Core Logic)
```go
func MergeSSTables(sstables []SSTable, output string) error {
    // 1. Open all inputs, read in sorted order
    // 2. For each key, keep only newest value
    // 3. Skip tombstones (deleted entries)
    // 4. Write merged result to output SSTable
    // 5. Update manifest, delete old files
}
```

---

## Interview Day Checklist

- [ ] Get good sleep night before
- [ ] Eat breakfast (sugar helps with talking)
- [ ] Practice your opening story once more
- [ ] Bring laptop ready to share screen
- [ ] Have code repo available
- [ ] Remember: They want to hire you (be confident)
- [ ] If asked something you don't know, say "I would investigate by..."

---

## Last Thoughts

**Most important:** Show you understand the "why" not just the "what"

- Why LSM? (write optimization)
- Why memtable? (fast batching)
- Why compaction? (manage file growth)
- Why WAL? (durability)

**Second most important:** Show you can debug and problem-solve

- "I hit issue X and solved it by..."
- "I optimized Y by profiling and finding..."
- "I tested for failure Z by..."

**Third most important:** Show you know the trade-offs

- "I chose to optimize for writes over reads because..."
- "I didn't implement feature X because it would..."
- "In production, I would..."

You've got this. You built something real and learned deeply. That's what matters.

---

## Before You Walk Into That Interview

Remember:
1. You understand LSM Trees deeply
2. You built a working database
3. You tested it thoroughly
4. You can explain every part
5. You know the trade-offs and limitations
6. You have real debugging stories

That's better than 90% of engineers. Go show them what you've got.

