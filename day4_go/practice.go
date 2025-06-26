package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

var mu sync.RWMutex

var books = make(map[string]Book)

func AddBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books[book.ID] = book
	c.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not that book"})
		return
	}
	delete(books, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Request Body",
			"details": err.Error(),
		})
		return
	}

	if _, exists := books[id]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not that book"})
		return
	}
	books[id] = book
	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
		"data":    book,
	})
}

func SearchAllBook(c *gin.Context) {
	bookList := make([]Book, 0, len(books))
	for _, book := range books {
		bookList = append(bookList, book)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Num":    len(bookList),
		"data":   bookList,
	})

}

func main() {
	r := gin.Default()
	r.POST("/book", AddBook)
	r.DELETE("/book/:id", DeleteBook)
	r.PUT("/book/:id", UpdateBook)
	r.GET("/books", SearchAllBook)

	mu.Lock()
	books["1"] = Book{ID: "1", Title: "Go语言编程", Author: "许式伟", Stock: "10"}
	books["2"] = Book{ID: "2", Title: "Clean Code", Author: "Robert Martin", Stock: "5"}
	mu.Unlock()

	r.Run(":8080")
}
