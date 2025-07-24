package router

import (
	"crud/api/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// r representa el router principal de Gin, es decir, el objeto que maneja todas las solicitudes (post, delete..)
func InitRoutes(r *gin.Engine, db *sql.DB) {
	//r *gin.Engine: el motor de rutas de Gin (tu servidor web)
	//db *sql.DB:  conexión a la base de datos
	// cuando se haga una solciitud get a books se va ejecutar una función que llama a los controlers
	r.GET("/books", func(c *gin.Context) {
		controllers.GetBooks(c, db)
	})

	// solicitud para POST
	r.POST("/books", func(c *gin.Context) {
		controllers.CreateBook(c, db)
	})
	// put
	r.PATCH("/books/:id", func(c *gin.Context) {
		controllers.UpdateBook(c, db)
	})
}
