package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/imageupload/routes"
	"github.com/olahol/go-imageupload"
)

// Variável global para armazenar a imagem atualmente carregada.
var currentImage *imageupload.Image

func main() {
	service := gin.Default()
	routes.Routes(service)
	service.Run(":8080")
}
