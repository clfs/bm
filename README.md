# bm
Workspace for quick Go benchmarks.

```text
go test -bench . -benchmem
```

```text
goos: darwin
goarch: arm64
pkg: github.com/clfs/bm
BenchmarkIsEmailGlobalRegexp-10     	33802567	        35.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsEmailChainedRegexp-10    	  397226	      2918 ns/op	    5729 B/op	      65 allocs/op
BenchmarkIsEmailVarRegexp-10        	  401308	      2934 ns/op	    5734 B/op	      65 allocs/op
PASS
ok  	github.com/clfs/bm	4.687s
```