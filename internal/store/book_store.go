package store

import (
	"api-rest/internal/model"
	"database/sql"
)

// Store es una interfaz que define los métodos para interactuar con el almacenamiento de datos.
type Store interface {
	GetAll() ([]*model.Libro, error)                         // GetAll devuelve una lista de todos los libros almacenados.
	GetByID(id int) (*model.Libro, error)                    // GetByID devuelve un libro específico por su ID.
	Create(libro *model.Libro) (*model.Libro, error)         // Create agrega un nuevo libro al almacenamiento.
	Update(id int, libro *model.Libro) (*model.Libro, error) // Update modifica un libro existente por su ID.
	Delete(id int) error                                     // Delete elimina un libro del almacenamiento por su ID.
}

// store es una implementación concreta de la interfaz Store que utiliza una base de datos SQL para almacenar los libros.
type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

// GetAll devuelve una lista de todos los libros almacenados en la base de datos.
func (s *store) GetAll() ([]*model.Libro, error) {
	q := `SELECT id, titulo, autor FROM libros`

	rows, err := s.db.Query(q)

	if err != nil {
		return nil, err
	}

	// Aseguramos que las filas se cierren después de procesarlas para liberar recursos.
	defer rows.Close()

	// Creamos una variable para almacenar los libros que se recuperarán de la base de datos.
	var libros []*model.Libro

	// Iteramos sobre las filas devueltas por la consulta y escaneamos cada fila en una estructura Libro.
	for rows.Next() {
		b := model.Libro{}
		if err := rows.Scan(&b.ID, &b.Titulo, &b.Autor); err != nil {
			return nil, err
		}
		libros = append(libros, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return libros, nil
}

// GetByID devuelve un libro específico por su ID desde la base de datos.
func (s *store) GetByID(id int) (*model.Libro, error) {
	q := `SELECT id, titulo, autor FROM libros WHERE id = ?`

	b := model.Libro{}
	err := s.db.QueryRow(q, id).Scan(&b.ID, &b.Titulo, &b.Autor)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// Create agrega un nuevo libro al almacenamiento de la base de datos.
func (s *store) Create(libro *model.Libro) (*model.Libro, error) {
	q := `INSERT INTO libros (titulo, autor) VALUES (?, ?)`

	// Ejecutamos la consulta de inserción y obtenemos el resultado para obtener el ID del libro recién insertado.
	result, err := s.db.Exec(q, libro.Titulo, libro.Autor)

	if err != nil {
		return nil, err
	}

	// Obtenemos el ID del libro recién insertado para asignarlo a la estructura Libro.
	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	// Asignamos el ID obtenido al libro y lo devolvemos.
	libro.ID = int(id)
	// Devolvemos el libro recién creado con su ID asignado.
	return libro, nil
}

// Update modifica un libro existente por su ID en la base de datos.
func (s *store) Update(id int, libro *model.Libro) (*model.Libro, error) {
	q := `UPDATE libros SET titulo = ?, autor = ? WHERE id = ?`

	_, err := s.db.Exec(q, libro.Titulo, libro.Autor, id)
	if err != nil {
		return nil, err
	}

	libro.ID = id
	return libro, nil
}

// Delete elimina un libro del almacenamiento por su ID en la base de datos.
func (s *store) Delete(id int) error {
	q := `DELETE FROM libros WHERE id = ?`

	_, err := s.db.Exec(q, id)
	return err
}
