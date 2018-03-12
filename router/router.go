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

	r.LoadHTMLGlob("view/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/todos", func(c *gin.Context){
		c.HTML(http.StatusOK,"todo.html",nil)
	})

	api := r.Group("")
	apiRouter(api)

	return r
}