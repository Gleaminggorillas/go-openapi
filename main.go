package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// configure the Gin server
	router := gin.Default()
	router.GET("/todo", getAllTodos)
	// run the Gin server
	router.Run()
}

type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

type message struct {
	Message string `json:"message"`
}

var todoList = []todo{
	{"1", "Learn Go"},
	{"2", "Build an API with Go"},
	{"3", "Document the API with swag"},
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}