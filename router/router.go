package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getrouter() *gin.Engine {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/fonts", "./public/fonts")

	r.LoadHTMLFiles("view/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})


	api := r.Group("")
	apiRouter(api)

	return r
}