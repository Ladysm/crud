package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// se realiza la estructura del libro con las propiedades que va tener
// Modelo simple del libro
type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublishedYear   int    `json:"publishedYear"`
	Genre           string `json:"genre"`
	ISBN            string `json:"isbn"`
	PageCount       int    `json:"pageCount"`
	Language        string `json:"language"`
	AvailableCopies int    `json:"availableCopies"`
}

// se usa el c (objeto de contexto)
// db *sql.DB la conexión a la base de datos para poder hacer consultas

func GetBooks(c *gin.Context, db *sql.DB) {
	// consulta SQL
	rows, err := db.Query("SELECT id, title, author, publishedYear, genre, isbn, pageCount, language, availableCopies FROM books")
	//manejo de error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar los libros"})
		return
	}
	// se cierra la coenxión
	defer rows.Close()
	// uso el struct, para declarar variable books que es una arreglo dinámico
	var books []Book
	// se itera por cada fila devuelta z la consulta
	for rows.Next() {
		//Por cada fila se crea una variable temporal b de tipo Book
		var b Book
		// lee la fila actual que viene de la base de datos y asigna cada columna a una variable en Go, si hay error lo devuelve
		if err := rows.Scan(
			&b.ID, &b.Title, &b.Author, &b.PublishedYear, &b.Genre, &b.ISBN, &b.PageCount, &b.Language, &b.AvailableCopies,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer los datos"})
			return
		}
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)
}
func CreateBook( c *gin.Context, db *sql.DB){
	var newBook
	if err := c.ShouldBindBodyWithJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Json invalido"})
	}
}