package model

// Libro representa un libro con su ID, título y autor.
type Libro struct {
	//ID     int    `xml:"id"`
	ID     int    `json:"id"`     // ID es un identificador único para cada libro
	Titulo string `json:"titulo"` // Titulo es el título del libro
	Autor  string `json:"autor"`  // Autor es el nombre del autor del libro
}
