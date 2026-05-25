# Go testing

Projecto simple con ejemplos de como hacer testing en Go.

## Descripción

El proyecto consiste en un API REST `/reverse[?q=STRING]`, que devuelve la cadena que se pase en el parámetro `q` invertido. Si no se pasa parámetro devuelve la cadema **Hola Mundo!** dada la vuelta.

## Testing

Se hacen varios tipos de test:

- Unitarios
- Integración
- Rendimiento(_benchmarking_)
- Fuzzy

### Unitarios

```
go test ./...
```

### Integración

### Rendimiento

```bash
$ go test -bench=. -benchmem

goos: linux
goarch: amd64
pkg: github.com/jomoespe/go-testing
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkReverse-16    	74837330	        16.00 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/jomoespe/go-testing	1.209s
```

La primera columna es el número de iteraciones (74837330). La segunda es el tiempo promedio de ejecución (16.00 ns/op) en manosegundos. La tercera (0 B/op) es el numero de bytes alocados por operación.

### Fuzzy

## See also

- [Benchmark Testing in Go: A Practical Guide](https://medium.com/@debug-ing/benchmark-testing-in-go-a-practical-guide-2900e008ce43)
