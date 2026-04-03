package service

import (
	"api-rest/internal/model"
	"api-rest/internal/store"
	"errors"
)

// Service es la estructura que representa el servicio de libros, que interactúa con la tienda para obtener los datos.
type Service struct {
	store store.Store
}

// New crea una nueva instancia del servicio de libros con la tienda proporcionada.
func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

// ObtenerTodosLosLibros devuelve una lista de todos los libros disponibles en la tienda.
func (s *Service) ObtenerTodosLosLibros() ([]*model.Libro, error) {
	libros, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}
	return libros, nil
}

// ObtenerLibroPorID devuelve un libro específico basado en su ID, o un error si no se encuentra.
func (s *Service) ObtenerLibroPorID(id int) (*model.Libro, error) {
	return s.store.GetByID(id)
}

// CrearLibro agrega un nuevo libro a la tienda después de validar que los campos necesarios estén presentes.
func (s *Service) CrearLibro(libro model.Libro) (*model.Libro, error) {
	// Validación básica para asegurarse de que se proporcionen los campos necesarios
	if libro.Titulo == "" || libro.Autor == "" {
		return nil, errors.New("el título y el autor son campos obligatorios") // Validación básica para asegurarse de que se proporcionen los campos necesarios
	}
	// Llamar al método Create de la tienda para agregar el nuevo libro
	return s.store.Create(&libro)
}

// ActualizarLibro actualiza un libro existente en la tienda basado en su ID, o devuelve un error si no se encuentra.
func (s *Service) ActualizarLibro(id int, libro model.Libro) (*model.Libro, error) {
	return s.store.Update(id, &libro)
}

// EliminarLibro elimina un libro existente en la tienda basado en su ID, o devuelve un error si no se encuentra.
func (s *Service) RemoverLibro(id int) error {
	return s.store.Delete(id)
}
