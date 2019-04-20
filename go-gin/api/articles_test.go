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

		envelop := ArticleEnvelope{
			Article: Article{
				Title:       "How to train your dragon",
				Description: "Ever wonder how?",
			}}

		Convey("Posting a new Article", func() {
			response := postJSON(api, "/articles/", envelop)

			So(response.Code, ShouldEqual, 201)

			Convey("Response", func() {
				var r ArticleEnvelope

				err := json.Unmarshal(response.Body.Bytes(), &r)

				So(err, ShouldBeNil)
				So(r, ShouldResemble, envelop)
			})
		})
	})
}
