package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Libro struct {
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Cantidad int    `json:"cantidad"`
}

// Para detectar si "cantidad" vino ausente o en 0
type libroIn struct {
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Cantidad *int   `json:"cantidad"` // puntero para saber si vino el campo
}

var (
	inventario = []Libro{
		{Titulo: "Cien años de soledad", Autor: "Gabriel García Márquez", Cantidad: 3},
		{Titulo: "Don Quijote de la Mancha", Autor: "Miguel de Cervantes", Cantidad: 2},
		{Titulo: "El Principito", Autor: "Antoine de Saint-Exupéry", Cantidad: 5},
		{Titulo: "1984", Autor: "George Orwell", Cantidad: 4},
		{Titulo: "Rayuela", Autor: "Julio Cortázar", Cantidad: 1},
	}
	mu sync.Mutex // para concurrencia segura
)

func main() {
	http.HandleFunc("/libros", librosHandler)
	log.Println("Servidor escuchando en http://127.0.0.1:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func librosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarLibros(w, r)
	case http.MethodPost:
		agregarLibro(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func listarLibros(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(inventario)
}

func agregarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var in libroIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if in.Titulo == "" || in.Autor == "" {
		http.Error(w, `{"error":"Faltan campos requeridos: 'titulo' y 'autor'."}`, http.StatusBadRequest)
		return
	}

	cantidad := 1
	if in.Cantidad != nil {
		cantidad = *in.Cantidad
	}
	if cantidad < 1 {
		http.Error(w, `{"error":"'cantidad' debe ser un entero >= 1."}`, http.StatusBadRequest)
		return
	}

	libro := Libro{Titulo: in.Titulo, Autor: in.Autor, Cantidad: cantidad}

	mu.Lock()
	inventario = append(inventario, libro)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": "Libro agregado",
		"libro":   libro,
	})
}
