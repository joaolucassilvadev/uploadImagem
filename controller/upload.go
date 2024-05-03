package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/krishpranav/imageupload/service"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func UploadImg(ctx *gin.Context) {
	// Processa a imagem enviada na requisição.
	img, err := imageupload.Process(ctx.Request, "file")
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define o caminho para a pasta onde você deseja salvar a imagem.
	// Certifique-se de que a pasta existe e que você tem permissões de gravação.
	targetFolder := "/home/joao/uploadImagem/uploadImagem/img"
	img.Filename = "1.png"
	filePath := filepath.Join(targetFolder, img.Filename)

	err = img.Save(filePath)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	currentImage = img
	service.Envio(filePath)

	fmt.Println("linkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk:")
	fmt.Print(filePath)
	os.Remove(filePath)

	ctx.Redirect(http.StatusMovedPermanently, "/")
}
