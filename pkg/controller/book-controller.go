package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-gorm-docker-books-api/pkg/model"
)

var Book model.Book

func GetBooksHandler(ctx *gin.Context) {
	books := model.GetBooks()
	ctx.JSON(http.StatusOK, books)
}

func GetBookByISBNHandler(ctx *gin.Context) {
	isbn := ctx.Param("isbn")
	bookISBN, err := strconv.ParseInt(isbn, 10, 64)
	if err != nil {
		log.Printf("error parsing ISBN: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ISBN"})
		return
	}

	book, _ := model.GetBookByISBN(bookISBN)
	ctx.JSON(http.StatusOK, book)
}

func CreateBookHandler(ctx *gin.Context) {
	var newBook model.Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		log.Printf("error parsing request body: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	book := newBook.CreateBook()
	ctx.JSON(http.StatusOK, book)
}

func DeleteBookHandler(ctx *gin.Context) {
	isbn := ctx.Param("isbn")
	bookISBN, err := strconv.ParseInt(isbn, 10, 64)
	if err != nil {
		log.Printf("error parsing ISBN: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ISBN"})
		return
	}

	model.DeleteBook(bookISBN)
	ctx.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

func UpdateBookHandler(ctx *gin.Context) {
	isbn := ctx.Param("isbn")
	bookISBN, err := strconv.ParseInt(isbn, 10, 64)
	if err != nil {
		log.Printf("error parsing ISBN: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ISBN"})
		return
	}

	var updatedBook model.Book
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		log.Printf("error parsing request body: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	book, _ := model.GetBookByISBN(bookISBN)
	if book == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	book.UpdateBook(updatedBook)
	ctx.JSON(http.StatusOK, book)
}
