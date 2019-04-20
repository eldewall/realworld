package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func postJSON(r http.Handler, path string, data interface{}) *httptest.ResponseRecorder {
	b, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(b))
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	return rec
}

func TestArticlesAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Convey("Creating an article", t, func() {
		api := gin.Default()

		Use(api)

		Convey("Posting a new Article", func() {
			envelop := ArticleEnvelope{
				Article: Article{
					Title:       "How to train your dragon",
					Description: "Ever wonder how?",
				}}
			response := postJSON(api, "/articles", envelop)

			So(response.Code, ShouldEqual, 201)
			So(response.Body.String(), ShouldEqual, "Foo")
		})
	})
}
