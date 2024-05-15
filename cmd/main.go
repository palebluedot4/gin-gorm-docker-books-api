package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gin-gorm-docker-books-api/pkg/config"
	"gin-gorm-docker-books-api/pkg/controller"
	"gin-gorm-docker-books-api/pkg/model"
)

func main() {
	router := gin.Default()
	config.ConnectDB()
	db := config.GetDB()
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	router.GET("/books", controller.GetBooksHandler)
	router.GET("/books/:isbn", controller.GetBookByISBNHandler)
	router.POST("/books", controller.CreateBookHandler)
	router.DELETE("/books/:isbn", controller.DeleteBookHandler)
	router.PUT("/books/:isbn", controller.UpdateBookHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
