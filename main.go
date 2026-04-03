package main

import (
	"api-rest/internal/service"
	"api-rest/internal/store"
	"api-rest/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {

	// Conectar a la base de datos
	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// Crear el table si no existe
	q := `
	CREATE TABLE IF NOT EXISTS libros (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo TEXT NOT NULL,
		autor TEXT NOT NULL
	);
	`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inyectar nuestras dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	// Configurar rutas
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Servidor ejecutandose en http://localhost:8080")
	fmt.Println(" API Endpoints:")
	fmt.Println(" GET    /books      - Obtener todos los libros")
	fmt.Println(" GET    /books/{id} - Obtener un libro por su ID")
	fmt.Println(" POST   /books      - Crear un nuevo libro")
	fmt.Println(" PUT    /books/{id} - Actualizar un libro existente")
	fmt.Println(" DELETE /books/{id} - Eliminar un libro existente")

	// Empezar y escuchar al servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
}
