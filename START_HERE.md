# START HERE - 3-5 Day LSM Database Crash Course

You have **3-5 days** to build a working LSM Tree database and ace your interview.

## 🚀 What To Do Right Now

1. **Open:** `CRASH_COURSE_3_5_DAYS.md`
2. **Read it completely** (takes ~20 minutes)
3. **Start Day 1** today
4. **Code every day** for the next 3-5 days

That's it. You don't need anything else.

---

## 📋 The Plan At A Glance

| Day | Focus | Hours | Deliverable |
|-----|-------|-------|-------------|
| **Day 1** | Learn LSM + Design | 6-8 | Understand architecture |
| **Day 2** | MemTable + WAL + SSTable | 8-10 | Three working components |
| **Day 3** | Database + Compaction | 8-10 | Full working database |
| **Day 4** | Testing + Polish | 6-8 | Clean code + tests |
| **Day 5** | Interview prep | 4-6 | Mastered 8 Q&A |

---

## ✅ What You'll Have At The End

✓ Working LSM Tree database  
✓ All code from scratch (no AI help)  
✓ Deep understanding of how it works  
✓ 8 perfect interview answers memorized  
✓ Confidence in technical discussions  

---

## 📖 Two Files Only

### 1. CRASH_COURSE_3_5_DAYS.md (11 KB)
**Your complete implementation roadmap**
- What to build each day
- Code snippets for guidance (not copy-paste)
- 8 key interview questions with perfect answers
- How to practice your pitch

Read this now.

### 2. QUICK_REFERENCE.md (11 KB)
**Your interview cheat sheet**
- 60-second pitch
- 5-minute explanation
- Common questions and answers
- Code snippets to memorize

Study this 1-2 days before interview.

---

## 🎯 Daily Breakdown

### Day 1: Understand The Problem (6-8 hours)
- Learn what LSM Tree is
- Understand why it's fast for writes
- Design the 5 components
- Setup Go project
- **Goal:** Can draw architecture on paper

### Day 2: Build The Storage Layer (8-10 hours)
- Implement MemTable (in-memory)
- Implement WAL (write-ahead log)
- Implement SSTable (disk storage)
- Test each independently
- **Goal:** Can put/get data

### Day 3: Glue It All Together (8-10 hours)
- Implement Database orchestrator
- Implement Compaction
- Implement recovery
- Integration test everything
- **Goal:** Full working database

### Day 4: Polish & Test (6-8 hours)
- Write comprehensive tests
- Stress test (10,000+ operations)
- Document your code
- Write README
- **Goal:** Production-ready code

### Day 5: Master Interview (4-6 hours)
- Memorize 8 key Q&A
- Practice 60-second pitch
- Explain each component
- Record yourself
- **Goal:** Interview ready

---

## 🔥 Quick Code Reference

### MemTable (In-Memory)
```go
type MemTable struct {
    data map[string]string
}

Put(key, value)    // O(1) - fast
Get(key)           // O(1) - fast
Delete(key)        // O(1) - fast
Entries()          // returns sorted entries
```

### WAL (Write-Ahead Log)
```go
type WAL struct {
    file *os.File
}

Append(entry)      // write to disk (durable)
Read()             // read all entries (for recovery)
```

### SSTable (Disk Storage)
```go
type SSTable struct {
    entries []Pair  // sorted
}

Write()            // serialize to file
Read(path)         // deserialize from file
Get(key)           // binary search
```

### Database (Orchestrator)
```go
type DB struct {
    memtable *MemTable
    wal *WAL
    sstables []*SSTable
}

Put(key, value)    // WAL + MemTable
Get(key)           // MemTable first, then SSTables
Compact()          // merge SSTables
```

---

## 💡 8 Key Interview Questions

1. **"What is an LSM Tree?"**
   - Answer: Write-optimized storage. Appends to memory/log, background compaction merges files.

2. **"Why LSM instead of B-Tree?"**
   - Answer: 100x faster writes (sequential vs random I/O). Trade-off: slower reads.

3. **"Walk me through a write."**
   - Answer: WAL → MemTable → flush when full → SSTable → compaction when many files.

4. **"Walk me through a read."**
   - Answer: Check MemTable first (O(1)), then SSTables (O(k log n)), return not found.

5. **"What if we crash?"**
   - Answer: WAL is durable. On recovery, replay WAL to restore memtable.

6. **"Why immutable SSTables?"**
   - Answer: Simplifies concurrency, enables efficient merging, allows background ops.

7. **"How does compaction work?"**
   - Answer: Merge multiple old SSTables into one. Reduces files, improves reads.

8. **"What's the read amplification?"**
   - Answer: Number of SSTables to check. ~10 in typical setup. Solved with bloom filters.

Memorize these 8 answers word-for-word.

---

## 📊 Success Criteria

**Day 1 ✅**
- Can draw LSM architecture from memory
- Can explain write path to someone
- Project setup complete

**Day 2 ✅**
- MemTable passes: put, get, delete, size tests
- WAL passes: append, read, persistence tests
- SSTable passes: write, read, search tests

**Day 3 ✅**
- Database Put/Get/Delete works
- Compaction triggers and merges correctly
- Stress test: 10,000 operations complete
- Data persists after close/reopen

**Day 4 ✅**
- All tests pass
- Code is clean and commented
- README documents your implementation
- Zero compiler warnings

**Day 5 ✅**
- Can recite 8 Q&A answers
- Practiced pitch 10+ times
- Can walk through code line-by-line
- Feel confident

---

## ⚠️ Common Mistakes (Avoid These)

❌ Overthinking the design  
✅ Simple design, iterate quickly

❌ Writing perfect code first time  
✅ Make it work first, optimize later

❌ Not testing as you go  
✅ Test each component immediately

❌ Trying to implement everything at once  
✅ Implement sequentially: MemTable → WAL → SSTable → DB

❌ Memorizing without understanding  
✅ Understand the "why" first, then memorize

---

## 📞 If You Get Stuck

1. **Stuck on theory?** Read CRASH_COURSE_3_5_DAYS.md Day 1 again
2. **Stuck on MemTable?** It's just a map with Put/Get/Delete
3. **Stuck on WAL?** It's just append-only file writes
4. **Stuck on SSTable?** Write sorted data to file, read back
5. **Stuck on Database?** Orchestrate the three components

If still stuck: Research similar implementations (RocksDB, LevelDB). Don't copy code, understand concepts.

---

## 🎤 Interview Pitch Practice

**60-Second Version:**
> "I built an LSM Tree database to understand write-optimized storage. LSM Trees achieve fast writes by batching data in memory and appending sequentially instead of random disk I/O. When the in-memory table fills, it flushes to a disk file (SSTable). Reads are slower because they check multiple files, but background compaction merges files to improve performance. The key insight is using immutable files and write-ahead logs for durability while maintaining high write throughput."

Practice this 10 times. Record yourself. Does it sound natural?

---

## 🏁 Final Checklist

**Before you start:**
- [ ] Have 3-5 days completely blocked (no other commitments)
- [ ] Have Go installed and working
- [ ] Have text editor ready (VS Code, etc)
- [ ] Have quiet place to work

**Day by day:**
- [ ] Completed Day 1 (understand LSM)
- [ ] Completed Day 2 (MemTable, WAL, SSTable)
- [ ] Completed Day 3 (Database, Compaction)
- [ ] Completed Day 4 (tests, docs)
- [ ] Completed Day 5 (interview prep)

**Before interview:**
- [ ] Code compiles without warnings
- [ ] All tests pass
- [ ] Can explain every line of code
- [ ] Practiced pitch 10+ times
- [ ] Memorized 8 Q&A answers
- [ ] Feel confident

---

## 🚀 You've Got This

You're building something real. In 3-5 days, you'll have:
- A working database
- Deep systems understanding
- Interview confidence
- Proof of your abilities

Every engineer who succeeded did this exact thing.

**Start now. Open CRASH_COURSE_3_5_DAYS.md.**

Time is ticking. Let's go! 💪
