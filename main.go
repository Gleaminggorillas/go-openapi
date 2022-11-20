package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// configure the Gin server
	router := gin.Default()

	// run the Gin server
	router.Run()
}
