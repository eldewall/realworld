package main

import (
	api "conduit/api"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api.Use(server)
}
