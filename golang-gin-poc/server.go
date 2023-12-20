package main

import (
	"golang-gin-poc/controller"
	"golang-gin-poc/middlewares"
	"golang-gin-poc/repository"
	"golang-gin-poc/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	VideoService    service.VideoService       = service.New(videoRepository)
	VideoController controller.VideoController = controller.New(VideoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	defer videoRepository.CloseDB()

	setupLogOutput()

	server := gin.New()

	server.Static("css", "./templates/css")

	server.LoadHTMLGlob("template/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := VideoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "success!!"})
			}

		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := VideoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "success!!"})
			}

		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := VideoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "success!!"})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	server.Run(":8080")
}
