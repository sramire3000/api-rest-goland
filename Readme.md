# API CON GOLANG

### Instalar paquetes
```
go get modernc.org/sqlite

go mod tidy
```

### Alternativa (con cgo): go-sqlite3
Si prefieres usar github.com/mattn/go-sqlite3, necesitas habilitar cgo y tener un compilador C (gcc) instalado.

Ejemplo en Windows (PowerShell):
```
setx CGO_ENABLED 1
```

Luego instala un toolchain con gcc (por ejemplo, MSYS2/MinGW) y valida:
```
gcc --version
```

Despues, cambia el import/driver en el codigo a sqlite3.

### Ejecutar
```
go run main.go
```

### Troubleshooting
1) Error: Binary was compiled with CGO_ENABLED=0, go-sqlite3 requires cgo to work
- Causa: estas usando go-sqlite3 sin cgo habilitado.
- Solucion A (recomendada): usar modernc.org/sqlite (pure Go), como esta configurado en este proyecto.
- Solucion B: habilitar cgo e instalar gcc (MSYS2/MinGW).

2) Error: gcc no se reconoce
- Causa: no hay compilador C en el PATH.
- Solucion: instala MSYS2/MinGW y verifica con:
```
gcc --version
```

3) Error: listen tcp :8080: bind: Only one usage of each socket address...
- Causa: el puerto 8080 ya esta en uso.
- Solucion A: detener el proceso que usa ese puerto.
- Solucion B: cambiar el puerto en el codigo (por ejemplo, de :8080 a :8081).