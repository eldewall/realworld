package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Article ...
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ArticleEnvelope we want getting posted in
type ArticleEnvelope struct {
	Article Article `json:"article"`
}

func postArticleHandler(c *gin.Context) {
	var envelop ArticleEnvelope
	err := c.BindJSON(&envelop)

	if err != nil {
		log.Printf("Invalid request: %s", err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(201, envelop.Article)
	}
}

// Use applies routes to our articles group
func Use(engine *gin.Engine) {
	articles := engine.Group("/articles")
	{
		articles.POST("/", postArticleHandler)
	}
}
