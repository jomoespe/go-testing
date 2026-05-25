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

### Fuzzy

