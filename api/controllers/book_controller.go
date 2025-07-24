package controllers

import (
	"crud/api/models"
	"crud/api/validators"
	"strings"

	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// se realiza la estructura del libro con las propiedades que va tener
// Modelo simple del libro

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
	var books []models.Book
	// se itera por cada fila devuelta z la consulta
	for rows.Next() {
		//Por cada fila se crea una variable temporal b de tipo Book
		var b models.Book
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
	var newBook models.Book
	//ShouldBindJSON  se usa para guardar el input recibido en el body(input)
	// si el json tiene campos faltante so inválidos, se guardan en err
	if err := c.ShouldBindJSON(&newBook); err != nil {
		// si hay error se devuelve status code 400
		//400 (Bad Request) c
		c.JSON(http.StatusBadRequest, gin.H{"error": "Json invalido"})

	}

	if errors := validators.ValidateBook(newBook); len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
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
	c.JSON(http.StatusCreated, gin.H{
		"code": "ok"})
}

// metodo put
func UpdateBook(c *gin.Context, db *sql.DB) {
	var updatedBook models.Book
	// obtenog el id
	id := c.Param("id")
	// Decodifica el JSON con los campos que se desean actualizar

	// Decodifica el JSON con los campos que se desean actualizar
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	if errors := validators.ValidateBook(updatedBook); len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}
	// Construye la query dinámica
	setClauses := []string{}
	args := []interface{}{}
	for field, value := range updateData {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", field))
		args = append(args, value)
	}
	args = append(args, id)
	query := fmt.Sprintf("UPDATE books SET %s WHERE id = ?", strings.Join(setClauses, ", "))
	// Ejecuta la consulta
	_, err := db.Exec(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el libro"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Libro actualizado correctamente"})
}
