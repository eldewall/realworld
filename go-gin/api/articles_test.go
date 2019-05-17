package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticlesSuite struct {
	suite.Suite
	API *gin.Engine
}

func (s *ArticlesSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	s.API = gin.Default()

	Use(s.API)
}

func (s *ArticlesSuite) TestPost() {
	envelop := ArticleEnvelope{
		Article: Article{
			Title:       "How to train your dragon",
			Description: "Ever wonder how?",
		}}

	r := postJSON(s.API, "/articles/", envelop)

	var received ArticleEnvelope

	json.Unmarshal(r.Body.Bytes(), &received)

	assert.Equal(s.T(), envelop, received)
	assert.Equal(s.T(), r.Code, 201)
}

func postJSON(r http.Handler, path string, data interface{}) *httptest.ResponseRecorder {
	b, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(b))

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	return rec
}

func TestArticlesSuite(t *testing.T) {
	suite.Run(t, new(ArticlesSuite))
}
