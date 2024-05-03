package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func UploadImg(ctx *gin.Context) {
	img, err := imageupload.Process(ctx.Request, "file")
	if err != nil {
		// Se ocorrer um erro durante o processamento, entra em pânico.
		panic(err)
	}

	// Atualiza a variável global com a nova imagem.
	currentImage = img

	ctx.Redirect(http.StatusMovedPermanently, "/")
}
