package transport

import (
	"api-rest/internal/model"
	"api-rest/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		libros, err := h.service.ObtenerTodosLosLibros()
		if err != nil {
			http.Error(w, "Error al obtener los libros", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libros)

	case http.MethodPost:
		var libro model.Libro

		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, "Error al decodificar el libro", http.StatusBadRequest)
			return
		}

		created, err := h.service.CrearLibro(libro)

		if err != nil {
			http.Error(w, "Error al crear el libro", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Método no disponible", http.StatusMethodNotAllowed)
	}
}

func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		libro, err := h.service.ObtenerLibroPorID(id)
		if err != nil {
			http.Error(w, "Libro no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libro)

	case http.MethodPut:
		var libro model.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, "Error al decodificar el libro", http.StatusBadRequest)
			return
		}

		updated, err := h.service.ActualizarLibro(id, libro)
		if err != nil {
			http.Error(w, "Error al actualizar el libro", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		if err := h.service.RemoverLibro(id); err != nil {
			http.Error(w, "Error al eliminar el libro", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método no disponible", http.StatusMethodNotAllowed)
	}
}
