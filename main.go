package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krishpranav/imageupload/routes"
	"github.com/olahol/go-imageupload"
)

// Variável global para armazenar a imagem atualmente carregada.
var currentImage *imageupload.Image

func main() {
	service := gin.Default()
	routes.Routes(service)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3031"
	}
	if err := service.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
