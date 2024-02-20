package controllers

import (
	"net/http"

	"github.com/AbhishekReddyy/Go-library-Management/database"
	"github.com/AbhishekReddyy/Go-library-Management/models"
	"github.com/gin-gonic/gin"
)

func Dashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookCollection := database.BookData(database.Client, "books")

		var books []models.Book

		cursor, err := bookCollection.Find(c.Request.Context(), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(c.Request.Context())

		for cursor.Next(c.Request.Context()) {
			var book models.Book
			if err := cursor.Decode(&book); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			books = append(books, book)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}

func AddBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func NewUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func BorrowBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ReturnBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SearchUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
