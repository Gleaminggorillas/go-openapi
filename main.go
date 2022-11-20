package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// configure the Gin server
	router := gin.Default()
	router.GET("/todo", getAllTodos)
	router.GET("/todo/:id", getTodoByID)
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

func getTodoByID(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and return item with matching ID
	for _, todo := range todoList {
			if todo.ID == ID {
					c.JSON(http.StatusOK, todo)
					return
			}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}