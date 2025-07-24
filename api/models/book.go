package models

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
