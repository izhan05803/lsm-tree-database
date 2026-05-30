# Interview Checklist - Before You Walk In

## ✅ Have You Done This?

### Code Understanding
- [ ] Read memtable.go line by line (understand why each line exists)
- [ ] Read wal.go - understand critical flush on line 68
- [ ] Read db.go - understand Put/Get/Compact flow
- [ ] Traced through an example Put operation on paper
- [ ] Traced through a recovery scenario mentally
- [ ] Explained each component to yourself out loud

### Practice
- [ ] Can explain the 60-second pitch smoothly
- [ ] Can answer all 8 interview Q&As from memory
- [ ] Can draw the architecture (MemTable → WAL → SSTable)
- [ ] Can explain why WAL comes before MemTable
- [ ] Can explain why SSTables are checked in reverse order

### Technical Details
- [ ] Know the thresholds (20 for flush, 5 for compact)
- [ ] Understand why HashMap is used (O(1) operations)
- [ ] Understand why JSON for WAL, Gob for SSTable
- [ ] Know what tombstone means and why it's needed
- [ ] Can explain recovery process step-by-step

### Edge Cases
- [ ] What happens if crash during Flush? (SSTable partially written)
- [ ] What happens if update existing key? (size doesn't increase)
- [ ] What happens if delete then get? (returns error)
- [ ] What happens if restart after delete? (still deleted, WAL replayed)

---

## The 60-Second Pitch

Practice saying this smoothly (should take 45-60 seconds):

"I built an LSM Tree database in Go from scratch. It's a write-optimized key-value store with three core components:

First, a MemTable - a fast in-memory hash map where all writes go initially. Every write is first logged to a Write-Ahead Log on disk for durability, then written to MemTable.

When MemTable reaches 20 entries, it's flushed to disk as an SSTable - a sorted, immutable file. The sorting is critical for the LSM Tree design.

As more writes come in, multiple SSTables accumulate. To keep reads fast, when the count reaches 5 SSTables, I trigger compaction: merging the two oldest tables into one sorted file and deleting the originals.

For reads, we first check MemTable (O(1)), then check SSTables in newest-first order. If the system crashes, we replay the WAL to restore state.

All core operations - Put, Get, Delete - work correctly with full crash recovery. The design trades slightly slower reads for very fast writes, which is optimal for write-heavy workloads."

---

## Answer These Perfectly

### Q1: Why WAL before MemTable?
"If we put to MemTable first, a crash could leave data in memory but not on disk. When the system restarts, that data is lost. By writing WAL first, we guarantee the operation is on persistent storage. If a crash happens after WAL but before MemTable, we replay the WAL on recovery."

### Q2: What's the Read Path?
"MemTable first (O(1) lookup), then SSTables in reverse order. We check newest SSTables first because they have the most recent data. This ensures we get the latest version of any key."

### Q3: What Happens on Recovery?
"Read the entire WAL file and replay all operations into a fresh MemTable. Each PUT overwrites any previous value (harmless), each DELETE removes the key. After replaying all entries, MemTable is fully restored."

### Q4: Why Compact?
"Every Get that misses MemTable has to check multiple SSTables. With many files, this is slow. Compaction merges multiple files into fewer files, reducing the number of lookups needed per Get."

### Q5: Why HashMap for MemTable?
"O(1) Put/Get/Delete operations. The tradeoff is that reads from multiple SSTables are slower, but we optimize for fast writes. Most systems are write-heavy anyway."

### Q6: Size Tracking Purpose?
"The size counter tells us when MemTable has accumulated enough entries to justify flushing to disk. When size > 20, we flush."

### Q7: Tombstone Explanation?
"When deleting a key, we can't just remove it from SSTables (they're immutable). Instead, we store a tombstone value (<<DELETED>>). On reads, finding a tombstone means the key was explicitly deleted, not just missing."

### Q8: What Would You Improve?
"In production, I'd add: (1) bloom filters to skip SSTables without the key, (2) binary search within SSTables instead of linear scan, (3) compression for SSTables, (4) better compaction strategy (leveled vs tiered), (5) concurrent access with proper locking."

---

## Architecture Diagram (Draw This on Paper)

```
┌─────────────────────────────────────────────────────┐
│                   LSM TREE DATABASE                  │
├─────────────────────────────────────────────────────┤
│                                                       │
│  ┌─────────────────────────────────────────────┐   │
│  │            USER APPLICATION                 │   │
│  │   Put("key", "value") / Get("key")         │   │
│  └────────────────┬────────────────────────────┘   │
│                   │                                 │
│           ┌───────▼─────────┐                      │
│           │   Write Path    │                      │
│           │  (1) WAL.Append │                      │
│           │  (2) MemTable   │                      │
│           └───────┬─────────┘                      │
│                   │                                 │
│         ┌─────────▼──────────┐                     │
│         │  WAL File on Disk  │                     │
│         │ (Durable Recovery) │                     │
│         └────────────────────┘                     │
│                   │                                 │
│           ┌───────▼──────────────┐                │
│           │  MemTable (HashMap)  │                │
│           │  O(1) Put/Get/Delete │                │
│           │  Size: N entries     │                │
│           └───────┬──────────────┘                │
│                   │                                │
│          When Size > 20: FLUSH                    │
│                   │                                │
│        ┌──────────▼───────────────┐              │
│        │   SSTables on Disk       │              │
│        │  (Sorted, Immutable)     │              │
│        │  SSTable_0, _1, _2 ...   │              │
│        └──────────┬───────────────┘              │
│                   │                               │
│       When Count > 5: COMPACT                    │
│       (Merge oldest 2 into 1)                    │
│                   │                               │
│        ┌──────────▼───────────────┐              │
│        │   Read Path (Get)        │              │
│        │  1. Check MemTable (✓?)  │              │
│        │  2. Check SSTables(NEW→OLD)
│        │  Return first match      │              │
│        └─────────────────────────┘              │
│                                                   │
└─────────────────────────────────────────────────────┘
```

---

## File Checklist

✓ memtable.go - MemTable implementation
✓ wal.go - Write-Ahead Log
✓ sstable.go - Sorted String Table
✓ db.go - Database orchestrator
✓ main.go - Demo application
✓ STUDY_GUIDE.md - This detailed guide
✓ README.md - Project overview

---

## Last Minute Tips

1. **Don't memorize code**, understand the concepts
2. **Draw diagrams** when explaining (use paper/whiteboard)
3. **Explain trade-offs**: "We chose this because... the alternative would be..."
4. **Be honest about simplifications**: "In production, we'd use..."
5. **Show you understand the problem**: "The reason we need LSM Trees is..."
6. **Practice your explanation** 5-10 times before interview

---

## Red Flags to Avoid

❌ "I just copied code I found online"
❌ "I'm not sure why we do that"
❌ "I didn't know that part was important"
❌ "I can't explain the ordering"
❌ "I don't understand crash recovery"

✅ "I designed it this way because..."
✅ "The tradeoff is..."
✅ "In production, we'd..."
✅ "The critical insight is..."
✅ "Let me draw this for you..."

---

## YOU'RE READY

You have a working LSM Tree database. You have a study guide. You have interview answers. Now go practice explaining it clearly.

The interview isn't about perfect code - it's about showing you understand **why** each piece exists and **how** they fit together.

You've got this. 🚀
