# LSM Tree - Quick Study Guide for Interviews

## What You Need to Memorize

### 1. The Write Path (Most Important)
```
User calls: db.Put("key", "value")
    ↓
1. WAL.Append("PUT", key, value)     ← Write to disk FIRST (durable)
2. MemTable.Put(key, value)           ← Write to memory (fast)
3. Check: MemTable.Size() > 20?       ← Auto-flush?
4. If yes: MemTable.Flush() → SSTable ← Persists sorted data
5. Check: SSTable count > 5?          ← Auto-compact?
6. If yes: Compact()                  ← Merge old tables
```

**Why WAL first?** If crash after step 2, you lose data. If crash after step 1, you replay from WAL.

### 2. The Read Path
```
User calls: db.Get("key")
    ↓
1. Check MemTable (O(1) map lookup)   ← Fastest
2. If not found, check SSTables in REVERSE order (newest first)
3. Why reverse? Newest SSTable has latest value
4. Return first match found
```

**Why newest first?** Because if key exists in both SSTable_0 and SSTable_5, SSTable_5 is newer data.

### 3. The Recovery Path (After Crash)
```
System starts: NewDB()
    ↓
1. Open WAL file
2. Read all entries sequentially
3. For each entry:
   - If "PUT": call memtable.Put(key, value)
   - If "DELETE": call memtable.Delete(key)
4. After replay, MemTable has all data
5. SSTables on disk still exist from before crash
```

**Key insight:** Just replay the operations. Order doesn't matter for Put (overwrites are harmless), Delete works correctly.

---

## Components Explained Simply

### MemTable (memtable.go)
- **Data Structure:** HashMap (map[K]V)
- **Purpose:** Fast in-memory writes
- **Why HashMap?** O(1) Put/Get/Delete
- **Size tracking:** Counts entries to know when to flush
- **Flush:** Sorts all pairs, returns them, clears map

**Critical detail:** When Put overwrites a key, size doesn't increase. Size = unique keys, not operations.

### Write-Ahead Log (wal.go)
- **Format:** JSON, one entry per line
- **Example entry:** `{"Op":"PUT","Key":"alice","Value":"100"}`
- **Why JSON?** Human-readable for debugging, easy to parse
- **Why newline separator?** Easy to read line-by-line
- **Critical:** Flush to disk after EVERY write (line 68)
- **Mutex:** Prevents concurrent writes from corrupting file

**Key insight:** WAL is append-only. Old entries stay forever. On recovery, we just replay all of them.

### SSTable (sstable.go)
- **Data Structure:** File on disk containing sorted pairs
- **Encoding:** Gob (binary, faster than JSON)
- **Sorted:** Critical for fast searching and compaction
- **Tombstone:** `<<DELETED>>` marks deleted keys
- **Get:** Linear scan (simple) or binary search (production)

**Key insight:** SSTables are immutable. Once created, they're never modified. Compaction creates new ones.

### Database (db.go)
- **Orchestrates** Put/Get/Delete/Flush/Compact
- **Thresholds:**
  - Flush at 20 entries (prevent unbounded MemTable)
  - Compact at 5 SSTables (prevent unbounded reads)
- **Compaction:** Merges first 2 SSTables into 1
  - Reads all pairs from both
  - Creates merged MemTable
  - Writes new SSTable
  - Deletes old files
  - Reduces file count (speeds up reads)

---

## Interview Q&A - Your Script

### Q1: "Walk me through a Put operation"
A: "First, we write to the WAL for durability - if the system crashes after this, we can recover. Then we write to MemTable which is fast because it's just a HashMap. We check if MemTable size exceeded 20 entries, and if so, we flush it to an SSTable file. If SSTable count exceeds 5, we compact by merging the oldest two tables."

### Q2: "Why write to WAL BEFORE MemTable?"
A: "If we reverse the order, the system could crash after writing to MemTable but before WAL. When it restarts, the operation is lost because MemTable is in-memory and gets cleared. By writing WAL first, we guarantee that if the operation is in memory, it's also on disk."

### Q3: "How does recovery work?"
A: "When the database starts, it reads the entire WAL file and replays every operation into a fresh MemTable. Since we just replay sequentially, if a key was Put multiple times, the final Put wins. If a key was Deleted, the Delete removes it. After recovery, MemTable has the exact state from before the crash."

### Q4: "Why check SSTables in reverse order?"
A: "Because newer SSTables have more recent data. If a key exists in both an old and new SSTable, we want the new one. By iterating in reverse, we naturally find the newest version first."

### Q5: "What's the purpose of Flush?"
A: "When MemTable gets too large, it becomes a liability. Flush converts it to an SSTable (sorted file on disk), which is immutable and efficient. It also frees MemTable memory for new writes."

### Q6: "Why Compact?"
A: "Every Get operation has to check multiple SSTables if the key isn't in MemTable. With many SSTables, this is slow. Compact merges multiple files into fewer files, reducing the number of files we need to check."

### Q7: "What happens if you update a key that already exists?"
A: "In Put, we check if the key exists. If it does, we update the value but don't increment size (size only tracks unique keys). If it doesn't exist, we increment size. This prevents size from growing unbounded with repeated updates."

### Q8: "Why use Gob for SSTables but JSON for WAL?"
A: "WAL needs to be human-readable for debugging. If something goes wrong, we can open the file and see what operations happened. SSTables don't need to be readable - they just need to be fast. Gob is binary and faster than JSON, so it's better for performance."

---

## The BIG Picture

**LSM Tree = Write-Optimized Database**

```
Why write-optimized?
- Put goes to MemTable (fast, in-memory)
- WAL is append-only (sequential writes are fast)
- Flush and Compact are background operations

Tradeoff:
- Reads are slower because we check MemTable + multiple SSTables
- But in most systems, writes >> reads, so it's worth it

Example: Key-value store for logging. Millions of writes, few reads.
```

---

## Code Walkthrough - What to Know

### memtable.go - Line by Line
- Line 15-16: `data` = the HashMap, `size` = entry count
- Line 26-31: Put checks if key exists before incrementing size
- Line 35-42: Get returns (value, true) if found, (zero, false) if not
- Line 45-51: Delete only decrements size if key actually existed
- Line 59-70: Flush sorts pairs (O(n log n)), clears map

### wal.go - Critical Lines
- Line 28: `os.O_APPEND` means new writes go to end of file
- Line 51: `json.Marshal` converts entry to JSON bytes
- Line 68: `w.writer.Flush()` forces write to disk (CRITICAL)
- Line 42-43: `w.mu.Lock()` prevents concurrent writes (thread safety)

### db.go - Critical Logic
- Line 78-79: WAL write BEFORE memtable write (order matters!)
- Line 88-93: Flush triggers at 20 entries (automatic)
- Line 99-101: Check MemTable first (fast path)
- Line 105: Iterate SSTables in REVERSE (newest first)
- Line 152-157: Compact triggers at 5 SSTables

---

## Practice Questions to Ask Yourself

1. **Why can't we skip WAL?** (because then crashes lose data)
2. **Why not binary search in SSTables right now?** (simpler implementation, production would add it)
3. **What happens if compaction crashes halfway?** (old files still exist, could recover)
4. **Can we compact while Put is happening?** (current implementation no, but production would use locks)
5. **Why is tombstone needed?** (so Get knows key was deleted, not just missing)

---

## The Pitch (60 seconds)

"I built an LSM Tree database in Go from scratch. It's a write-optimized design: Puts go to a fast in-memory MemTable backed by a Write-Ahead Log for durability. When MemTable gets too large, it flushes to an SSTable on disk. Multiple SSTables are periodically compacted to reduce file count and speed up reads. All Put/Get/Delete operations work correctly, and the system fully recovers from crashes by replaying the WAL."

---

## Before Your Interview

1. Read memtable.go - understand why each line is there
2. Read wal.go - understand why flush is critical
3. Read db.go - understand the Put/Get/Compact paths
4. Run `go run .` several times - see the output
5. Answer the 8 Q&As from memory
6. Draw the architecture on paper (MemTable → WAL, SSTable, Compact flow)
7. Explain to someone else what each component does

You know this now. You just need to be able to explain it clearly.
