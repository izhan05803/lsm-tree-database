# Embedded LSM Tree Engine - 3-5 Day Crash Course

**Fast track version** for building an LSM Tree database in 3-5 days with zero AI assistance.

## ✅ Implementation Complete!

This repository now contains a **fully functional LSM Tree database** implementation in Go.

### What's Implemented:
- ✅ MemTable with hash-map storage (O(1) operations)
- ✅ Write-Ahead Log for durability and recovery
- ✅ SSTable for sorted on-disk storage
- ✅ Auto-flush when MemTable exceeds threshold (20 entries)
- ✅ Automatic compaction when SSTable count exceeds threshold (5 files)
- ✅ Complete WAL recovery on startup
- ✅ All put/get/delete operations working
- ✅ Comprehensive unit tests (10+ tests, all passing)
- ✅ Stress test with 1000+ operations

### Quick Run:
```bash
go run .                # Run demo
go test -v             # Run all tests
go build -o lsm-db.exe # Build binary
```

### Documentation:
- **IMPLEMENTATION.md** - Complete technical documentation of the implementation
- **CRASH_COURSE_3_5_DAYS.md** - Original 3-5 day learning guide
- **QUICK_REFERENCE.md** - Interview preparation guide

---

## 📚 Documentation Overview (Minimal & Focused)

### **CRASH_COURSE_3_5_DAYS.md** ⭐ START HERE
**Complete 3-5 day implementation guide** - Everything you need.

**Day 1:** Learn LSM theory + system design (6-8 hours)
**Day 2:** Implement MemTable + WAL + SSTable (8-10 hours)
**Day 3:** Implement Database core + Compaction (8-10 hours)
**Day 4:** Testing + Polish + Interview prep (6-8 hours)
**Day 5:** (Optional) Interview mastery (4-6 hours)

Contains:
- What to build each day
- Code snippets for each component
- 8 key interview questions with answers
- 60-second pitch practice
- Interview checklist

**This is your complete roadmap.** Follow day-by-day.

**Read time:** 20 minutes (then code)

---

### **QUICK_REFERENCE.md** (Interview Answers)
**Your interview script** - Key talking points.

- 60-second overview
- 5-minute deep dive explanation
- Common questions & perfect answers
- Code snippets to memorize
- Interview day checklist

**Memorize this section by section.**

**Read time:** 15 minutes (reference before interview)

---

## 🎯 How to Use This Repository (3-5 Days)

**Read CRASH_COURSE_3_5_DAYS.md now. It has everything you need.**

### Quick Start (Choose based on your timeline):

**If you have 5 days (8-10 hours/day):**
1. Read CRASH_COURSE_3_5_DAYS.md (20 min)
2. Follow Day 1-5 exactly as written
3. Study QUICK_REFERENCE.md before interview

**If you have 3 days (10-12 hours/day):**
1. Read CRASH_COURSE_3_5_DAYS.md (20 min)
2. Skip Day 4 testing, just make sure code works
3. Follow Days 1-3 + compress Day 5 (interview prep)
4. Study QUICK_REFERENCE.md intensively

**If you have 2 days (14-16 hours/day):**
1. Skim CRASH_COURSE_3_5_DAYS.md (10 min)
2. Do Day 1 morning only (2 hours theory)
3. Do Days 2-3 combined (implement everything fast)
4. Study QUICK_REFERENCE.md for interview
5. Focus on understanding 8 key Q&A

---

## 📂 Directory Structure

```
embedded-lsm-tree-engine/
├── 📖 Essential Documentation (Your Complete Path)
│   ├── README.md                         ← You are here
│   ├── CRASH_COURSE_3_5_DAYS.md         ← START HERE ⭐
│   └── QUICK_REFERENCE.md               ← Interview prep
│
└── 💻 Code Files (You Will Implement)
    ├── main.go                          ← Entry point
    ├── db.go                            ← Database core
    ├── memtable.go                      ← In-memory store
    ├── sstable.go                       ← Disk storage
    ├── wal.go                           ← Write-ahead log
    ├── manifest.go                      ← Metadata (optional)
    ├── compaction.go                    ← Background merging
    ├── go.mod                           ← Go module
    └── src/                             ← Optional organization
```
embedded-lsm-tree-engine/
├── README.md                          ← You are here
├── LSM_DB_REFERENCE.md               ← Architecture reference
├── REVERSE_ENGINEERING_GUIDE.md      ← Complete learning framework
├── WEEK_BY_WEEK_PLAN.md              ← Day-by-day schedule
├── IMPLEMENTATION_CHECKLIST.md       ← Task tracking
├── QUICK_REFERENCE.md                ← Interview talking points
├── go.mod                            ← Go module file
├── src/                              ← (Optional) Organization folder
├── main.go                           ← Entry point
├── db.go                             ← Database core (you write)
├── memtable.go                       ← In-memory store (you write)
├── sstable.go                        ← Sorted tables (you write)
├── compaction.go                     ← Background merging (you write)
├── wal.go                            ← Write-ahead log (you write)
├── manifest.go                       ← Metadata tracking (you write)
└── *_test.go                         ← Tests (you write)
```

---

## ⏱️ Timeline

- **3 days minimum:** Implement basic database, understand core concepts
- **4-5 days (recommended):** Implement everything, polish, interview prep
- **2 days (panic mode):** Focus on MemTable + Database + 8 interview Q&A

Total time: 24-40 hours (3-5 days × 8-10 hours/day)

---

## 🚀 Quick Start

### If you have only 2 weeks:
1. Skip deep theory, skim theory sections
2. Focus on: `WEEK_BY_WEEK_PLAN.md` Weeks 1-3
3. Implement fast, write tests
4. Polish code and interview prep

### If you have 4 weeks:
1. Follow `WEEK_BY_WEEK_PLAN.md` Weeks 1-4
2. Skip advanced optimizations
3. Focus on core functionality
4. Polish and prepare for interview

### If you have 8 weeks (ideal):
1. Follow everything sequentially
2. Deep understanding of all components
3. Thorough testing and performance tuning
4. Well-prepared for interview

---

## 📖 Reading Order (Fast Track)

**Today:**
1. **CRASH_COURSE_3_5_DAYS.md** - Read the entire thing (20 min)
2. Start **Day 1** immediately

**Next 3-5 days:**
- Follow **CRASH_COURSE_3_5_DAYS.md** day by day
- Code simultaneously as you read each day

**Before interview:**
- Study **QUICK_REFERENCE.md** intensively
- Practice the 8 key Q&A from crash course
- Run through your code once

---

## 🎓 Learning Outcomes

After completing this project, you will understand:

### System Design
- ✅ Trade-offs between write and read optimization
- ✅ How LSM Trees achieve high write throughput
- ✅ Durability guarantees (WAL, recovery)
- ✅ Background maintenance (compaction)
- ✅ Concurrency patterns in databases

### Implementation
- ✅ How to architect complex systems
- ✅ Binary file format design
- ✅ Serialization and deserialization
- ✅ Crash recovery mechanisms
- ✅ Performance optimization techniques

### Engineering Practices
- ✅ Test-driven development
- ✅ Debugging complex systems
- ✅ Performance profiling and tuning
- ✅ Code documentation and communication
- ✅ Problem-solving under constraints

### Interview Skills
- ✅ Explain architectural decisions clearly
- ✅ Tell compelling technical stories
- ✅ Answer questions about trade-offs
- ✅ Discuss optimization opportunities
- ✅ Defend design choices with reasoning

---

## 💡 Key Principles

### 1. **Understand the "Why"**
Don't just implement code. Understand why each component exists and what problem it solves.

### 2. **Build Bottom-Up**
Implement simple components first (MemTable), then build on top. This prevents debugging nightmares.

### 3. **Test Thoroughly**
Write tests as you go. This is how you catch bugs early and ensure correctness.

### 4. **Document Your Decisions**
Write down why you made each decision. This helps you explain it in interviews.

### 5. **Optimize Deliberately**
Only optimize after measuring. Use profiling to find real bottlenecks.

### 6. **Practice Your Explanation**
Your code isn't the project—your ability to explain it is. Practice talking through your work.

---

## 📝 Interview Tips

### Before Interview
- Practice explaining each component in 1-2 minutes
- Have 3-4 prepared stories about challenges you faced
- Know your performance numbers (throughput, latency)
- Be ready to discuss what you'd do differently

### During Interview
- Start with high-level architecture
- Drill down only when asked for details
- Be honest about what you simplified vs. production systems
- Show your problem-solving process, not just the solution
- Use analogies to explain complex concepts

### After Interview
- Ask what impressed them (feedback for future)
- Offer to discuss specific technical challenges
- Show enthusiasm for the problem domain

---

## 🤔 Common Questions Answered

### "Can I skip any documents?"
**Yes, but strategically:**
- Skip `QUICK_REFERENCE.md` if implementing before interview
- You need `LSM_DB_REFERENCE.md` for understanding structure
- `WEEK_BY_WEEK_PLAN.md` and `IMPLEMENTATION_CHECKLIST.md` are for tracking
- `REVERSE_ENGINEERING_GUIDE.md` is essential

### "How long will this really take?"
- Minimum: 30-40 hours if you're experienced with Go and databases
- Typical: 50-60 hours for most engineers
- Maximum: 70+ hours if learning Go or databases from scratch

### "Can I modify the architecture?"
**Yes!** That's the point. The guides show one approach. Consider alternatives:
- Different data structures for MemTable (skip list vs. map)
- Different compaction strategies (leveled vs. tiered)
- Different file formats
- Different recovery strategies

### "What if I get stuck?"
1. Read the relevant section of the guides again
2. Review the design documents you wrote
3. Check if simpler solution exists
4. Add logging/debugging
5. Search similar implementations (LevelDB, RocksDB)
6. Ask for help on conceptual issues, not code answers

---

## 📊 Success Metrics

You'll know you're ready for the interview when:

- [ ] Can draw LSM architecture from memory in 5 minutes
- [ ] Can explain each component's purpose in 1 minute
- [ ] Can walk through Put/Get/Compaction operations step-by-step
- [ ] Have told your origin story 10+ times smoothly
- [ ] All tests pass and code compiles without warnings
- [ ] Performance benchmarks documented and optimizations explained
- [ ] Can answer "why this design?" for each decision
- [ ] Have 2-3 stories about bugs you debugged and learned from
- [ ] Feel confident you understand the system deeply

---

## 🎯 Final Notes

**This is a challenging but rewarding project.** You're not just learning code—you're learning systems thinking. That's what separates good engineers from great engineers.

**Key mindset:**
- Embrace the struggle—it means you're learning
- Celebrate small wins—each component working is progress
- Document as you go—future you (and the interviewer) will thank you
- Teach others—explaining solidifies understanding

**Remember:** Every engineer who built a database went through this process. The difficulty is the point. You've got this.

---

## 📚 Additional Resources

### Papers
- LSM Tree original paper (O'Neil et al., 1996)
- RocksDB: A Persistent Key-Value Store for Flash and RAM Storage

### Implementations
- LevelDB (C++, reference implementation)
- RocksDB (Go/Java/Python bindings available)
- BadgerDB (Pure Go, modern approach)

### Blog Posts
- RocksDB design and optimization
- LevelDB documentation
- Crash course in LSM Trees

---

## 🚦 Status Checklist

Track your progress here:

- [ ] Read all documentation
- [ ] Completed Week 1 (foundation and design)
- [ ] Completed Week 2 (MemTable & WAL)
- [ ] Completed Week 3 (SSTable & Manifest)
- [ ] Completed Week 4 (Database core)
- [ ] Completed Week 5 (Compaction & Recovery)
- [ ] Completed Week 6 (Testing & polish)
- [ ] Started interview prep
- [ ] Done mock interview
- [ ] Ready for real interview!

---

**Good luck! You're building something real and learning deeply. That's what matters.**

Questions? Review the relevant guide section. Can't find the answer? Research the concept, design a solution, implement it, and learn.

Happy building! 🚀

