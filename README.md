# lsm-tree-database
LSM Tree database from scratch to understand how modern databases achieve high write throughput. LSM Trees optimize for writes by batching them in memory (memtable), then writing to disk sequentially. This is much faster than random disk I/O in B-Trees.
