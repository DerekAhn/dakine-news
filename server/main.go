package main

import (
	"github.com/gin-contrib/gzip"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/", index)
	router.Run(":3000")
}
