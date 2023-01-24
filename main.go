package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()
	// router.LoadHTMLGlob("templates/*.html")

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.html", gin.H{})
	// })

	// router.GET("/test", func(ctx *gin.Context) {

	// 	ctx.HTML(http.StatusOK, "test.html", gin.H{})
	// })
	// router.Run()
	engine:= gin.Default()
	engine.GET("/api/message", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
		})
	})
	engine.GET("/api/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
				"task": "hello world",
		})
	})
	engine.Run()
}