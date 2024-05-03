package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Image(ctx *gin.Context) {
	if currentImage == nil {
		// Se não houver imagem carregada, retorna um status HTTP 404 (Not Found).
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	// Escreve a imagem no corpo da resposta.
	currentImage.Write(ctx.Writer)
}
