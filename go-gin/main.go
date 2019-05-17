package main

import (
	api "conduit/api"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	a := api.New()

	api.Use(server)
}
