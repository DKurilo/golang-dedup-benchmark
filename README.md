# Benchmark for simple dedup implementation.

I almost sure, for simple deduplication implementation for strings, complexity can be more than O(N) in case we have mostly unique strings. So I tried to check it.

Results:  

```
bash-3.2$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkDedup/Dedup_justMap_2-8                10000000               171 ns/op
BenchmarkDedup/Dedup_preinitMap_2-8             10000000               178 ns/op
BenchmarkDedup/Dedup_justMap_4-8                 5000000               344 ns/op
BenchmarkDedup/Dedup_preinitMap_4-8              5000000               324 ns/op
BenchmarkDedup/Dedup_justMap_8-8                 3000000               583 ns/op
BenchmarkDedup/Dedup_preinitMap_8-8              2000000               623 ns/op
BenchmarkDedup/Dedup_justMap_16-8                1000000              2348 ns/op
BenchmarkDedup/Dedup_preinitMap_16-8             1000000              1577 ns/op
BenchmarkDedup/Dedup_justMap_32-8                 300000              4733 ns/op
BenchmarkDedup/Dedup_preinitMap_32-8              500000              2857 ns/op
BenchmarkDedup/Dedup_justMap_64-8                 200000             10226 ns/op
BenchmarkDedup/Dedup_preinitMap_64-8              300000              5574 ns/op
BenchmarkDedup/Dedup_justMap_128-8                100000             20512 ns/op
BenchmarkDedup/Dedup_preinitMap_128-8             200000             10749 ns/op
BenchmarkDedup/Dedup_justMap_256-8                 50000             38568 ns/op
BenchmarkDedup/Dedup_preinitMap_256-8             100000             21626 ns/op
BenchmarkDedup/Dedup_justMap_512-8                 20000             79994 ns/op
BenchmarkDedup/Dedup_preinitMap_512-8              50000             39558 ns/op
BenchmarkDedup/Dedup_justMap_1024-8                10000            150169 ns/op
BenchmarkDedup/Dedup_preinitMap_1024-8             20000             87264 ns/op
BenchmarkDedup/Dedup_justMap_2048-8                 3000            372551 ns/op
BenchmarkDedup/Dedup_preinitMap_2048-8             10000            180458 ns/op
BenchmarkDedup/Dedup_justMap_4096-8                 2000            714756 ns/op
BenchmarkDedup/Dedup_preinitMap_4096-8              5000            392516 ns/op
BenchmarkDedup/Dedup_justMap_8192-8                 1000           1606899 ns/op
BenchmarkDedup/Dedup_preinitMap_8192-8              2000            888833 ns/op
BenchmarkDedup/Dedup_justMap_16384-8                 500           3399580 ns/op
BenchmarkDedup/Dedup_preinitMap_16384-8             1000           2242878 ns/op
BenchmarkDedup/Dedup_justMap_32768-8                 200           8602999 ns/op
BenchmarkDedup/Dedup_preinitMap_32768-8              300           5270603 ns/op
BenchmarkDedup/Dedup_justMap_65536-8                 100          17904654 ns/op
BenchmarkDedup/Dedup_preinitMap_65536-8              200           9556238 ns/op
BenchmarkDedup/Dedup_justMap_131072-8                 30          39043726 ns/op
BenchmarkDedup/Dedup_preinitMap_131072-8             100          20814028 ns/op
BenchmarkDedup/Dedup_justMap_262144-8                 20          85120428 ns/op
BenchmarkDedup/Dedup_preinitMap_262144-8              30          45196552 ns/op
BenchmarkDedup/Dedup_justMap_524288-8                  5         220182494 ns/op
BenchmarkDedup/Dedup_preinitMap_524288-8              10         110840243 ns/op
BenchmarkDedup/Dedup_justMap_1048576-8                 2         515177287 ns/op
BenchmarkDedup/Dedup_preinitMap_1048576-8                      5         264400573 ns/op
BenchmarkDedup/Dedup_justMap_2097152-8                         1        1134440636 ns/op
BenchmarkDedup/Dedup_preinitMap_2097152-8                      2         569266185 ns/op
BenchmarkDedup/Dedup_justMap_4194304-8                         1        2176817763 ns/op
BenchmarkDedup/Dedup_preinitMap_4194304-8                      1        1075823090 ns/op
BenchmarkDedup/Dedup_justMap_8388608-8                         1        5170760459 ns/op
BenchmarkDedup/Dedup_preinitMap_8388608-8                      1        2209244710 ns/op
BenchmarkDedup/Dedup_justMap_16777216-8                        1        10951988496 ns/op
BenchmarkDedup/Dedup_preinitMap_16777216-8                     1        4634219900 ns/op
BenchmarkDedup/Dedup_justMap_33554432-8                        1        24239925962 ns/op
BenchmarkDedup/Dedup_preinitMap_33554432-8                     1        12572539593 ns/op
BenchmarkDedup/Dedup_justMap_67108864-8                        1        90527350257 ns/op
BenchmarkDedup/Dedup_preinitMap_67108864-8                     1        65869985182 ns/op
PASS
ok      .../benchgraph/dedupperf        495.863s
```

Visualization for Results are here:  
https://docs.google.com/spreadsheets/d/1OwH9gF4AokMkAZ1hDc5qi4w12POKDze6jlFHKlXDfoA/edit?usp=sharing  

## Conclusion

Yes, it's not O(N).
