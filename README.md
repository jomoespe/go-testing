# Go Testing

Here is a guide and a simple project structure to help you understand how testing works in Go.


> TODO add [vulnerabilities checking](https://go.dev/doc/tutorial/govulncheck)

## Project description and structure

The project is a simple `/reverse[?q=STRING]` REST API, that returns reverse `q` parameter, or reverse **Hello world!** string by default.

## Testing

This tests are done:

- Unit
- Integration [PENDING]
- Performance (_benchmarking_)
- Fuzzy

### Unit tests

> TBD describe

```bash
go test ./...
```

### Integration tests

> TBD describe

### Performance

> TBD describe

```bash
$ go test ./... -bench=. -benchmem

PASS
ok  	github.com/jomoespe/go-testing	0.005s
goos: linux
goarch: amd64
pkg: github.com/jomoespe/go-testing/pkg/stringutils
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkReverse-16    	19578816	        71.43 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/jomoespe/go-testing/pkg/stringutils	1.408s
```

First column is the number of iteratons, 19578816 in this case. Second column is the average execution time in nanos, 71.43 ns/op in this case. Third column is the number of alocated bytes by iteration, 8 B/op in this case. Last fourth column is the number of allocations done by iterations, 1 alloc/op in this case.

### Fuzzy

> TBD describe

```bash
go test ~/pkg/stringutils -fuzz=Fuzz -fuzztime 30s
```

## See also

- [Benchmark Testing in Go: A Practical Guide](https://medium.com/@debug-ing/benchmark-testing-in-go-a-practical-guide-2900e008ce43)
- [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
- [Coverage profiling support for integration tests](https://go.dev/doc/build-cover#glos-integration-test)
