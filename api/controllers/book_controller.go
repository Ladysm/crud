package controllers

import (
	"crud/api/validators"
	"database/sql"
	"fmt"
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

func CreateBook(c *gin.Context, db *sql.DB) {
	//se crea la variable y luego se asocia el objeto, esto es para guardar los datos del input body
	var newBook Book
	//ShouldBindJSON  se usa para guardar el input recibido en el body(input)
	// si el json tiene campos faltante so inválidos, se guardan en err
	if err := c.ShouldBindJSON(&newBook); err != nil {
		// si hay error se devuelve status code 400
		//400 (Bad Request) c
		c.JSON(http.StatusBadRequest, gin.H{"error": "Json invalido"})

	}

	// se valida los campos del input
	if !validators.IsValidGenre(newBook.Genre) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre invalid"})

	}
	// validador de language
	if validators.IsValidLanguage(newBook.Language) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "language invalid"})

	}
	if !validators.IsValidPublishedYear(newBook.PublishedYear) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Año de publicación inválido"})
		return
	}
	if !validators.IsValidPageCount(newBook.PageCount) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número de páginas inválido"})
		return
	}
	if !validators.IsValidAvailableCopies(newBook.AvailableCopies) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cantidad de copias inválida"})
		return
	}
	if !validators.IsValidISBN(newBook.ISBN) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ISBN inválido"})
		return
	}
	// Se define la consulta SQL pa insertar le libro
	query := `
		INSERT INTO books 
		(title, author, publishedYear, genre, isbn, pageCount, language, availableCopies)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	// se ejecuta el insert
	result, err := db.Exec(query,
		newBook.Title, newBook.Author, newBook.PublishedYear,
		newBook.Genre, newBook.ISBN, newBook.PageCount,
		newBook.Language, newBook.AvailableCopies,
	)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo insertar el libro"})
		return
	}
	// se obtiene el id del registro del libro creado
	insertedID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el ID insertado"})
		return
	}

	newBook.ID = int(insertedID)

	// se retorna el id del libro y el status code
	c.JSON(http.StatusCreated, newBook)
}
