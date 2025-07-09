package controllers

// se realiza la estructura del libro con las propiedades que va tener
// Modelo simple del libro
type Book struct {
	ID              int    `json:"id"`
	Author          string `json:"author"`
	PublishedYear   int    `json:"publishedYear"`
	Genre           string `json:"genre"`
	ISBN            string `json:"isbn"`
	PageCount       int    `json:"pageCount"`
	Language        string `json:"language"`
	AvailableCopies int    `json:"availableCopies"`
}
