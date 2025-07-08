package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/books", func(c *gin.Context) {
		controllers.GetBooks(c, db)
	})

}
