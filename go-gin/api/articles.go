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

// Store ...
type Store interface {
	Create(a *Article) error
}

type api struct {
	Store Store
}

// API unified
type API interface {
	Register(e *gin.Engine)
}

// New creates a Articles API
func New(s Store) API {
	return &api{Store: s}
}

func (api *api) Register(e *gin.Engine) {
	articles := e.Group("/articles")
	{
		articles.POST("/", api.onPost)
	}
}

func (api *api) onPost(c *gin.Context) {
	var envelop ArticleEnvelope
	err := c.BindJSON(&envelop)

	if err != nil {
		log.Printf("Invalid request: %s", err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(201, envelop)
	}
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
		c.JSON(201, envelop)
	}
}

// Use applies routes to our articles group
func Use(engine *gin.Engine) {
	articles := engine.Group("/articles")
	{
		articles.POST("/", postArticleHandler)
	}
}
