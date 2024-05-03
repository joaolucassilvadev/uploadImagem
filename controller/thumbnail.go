package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

func Thumbnail(ctx *gin.Context) {
	if currentImage == nil {
		// Se não houver imagem carregada, retorna um status HTTP 404 (Not Found).
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	// Gera uma miniatura JPEG da imagem com tamanho de 300x300 pixels e qualidade de 80%.
	t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)
	if err != nil {
		// Se ocorrer um erro durante o processo, entra em pânico.
		panic(err)
	}

	// Escreve a miniatura no corpo da resposta.
	t.Write(ctx.Writer)
}
