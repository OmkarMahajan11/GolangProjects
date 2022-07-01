package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"poc.com/gin-poc/controller"
	"poc.com/gin-poc/middleware"
)

var videoController controller.VideoController = controller.New()

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery())
	server.Use(middleware.Logger())
	server.Use(middleware.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":8080")
}
