package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

// Variável global para armazenar a imagem atualmente carregada.
var currentImage *imageupload.Image

func main() {
	// Inicializa um novo objeto Gin com as configurações padrão.
	r := gin.Default()

	// Define uma rota para o método GET em "/".
	// Quando um usuário acessa essa rota, o servidor envia o arquivo "index.html".
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	// Define uma rota para o método GET em "/image".
	// Esta rota serve para mostrar a imagem atualmente carregada.
	r.GET("/image", func(c *gin.Context) {
		if currentImage == nil {
			// Se não houver imagem carregada, retorna um status HTTP 404 (Not Found).
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		// Escreve a imagem no corpo da resposta.
		currentImage.Write(c.Writer)
	})

	// Define uma rota para o método GET em "/thumbnail".
	// Esta rota serve para gerar e exibir uma miniatura da imagem carregada.
	r.GET("/thumbnail", func(c *gin.Context) {
		if currentImage == nil {
			// Se não houver imagem carregada, retorna um status HTTP 404 (Not Found).
			c.AbortWithStatus(http.StatusNotFound)
		}

		// Gera uma miniatura JPEG da imagem com tamanho de 300x300 pixels e qualidade de 80%.
		t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)
		if err != nil {
			// Se ocorrer um erro durante o processo, entra em pânico.
			panic(err)
		}

		// Escreve a miniatura no corpo da resposta.
		t.Write(c.Writer)
	})

	// Define uma rota para o método POST em "/upload".
	// Esta rota é usada para fazer upload de uma nova imagem.
	r.POST("/upload", func(c *gin.Context) {
		// Processa a imagem enviada na requisição.
		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			// Se ocorrer um erro durante o processamento, entra em pânico.
			panic(err)
		}

		// Atualiza a variável global com a nova imagem.
		currentImage = img

		// Redireciona o usuário para a rota "/" após o upload bem-sucedido.
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Inicia o servidor na porta 8080, ouvindo e atendendo às requisições HTTP.
	r.Run(":8080")
}
