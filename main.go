package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"

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
			// Se ocorrer um erro durante o processamento, responde com o erro.
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Define o caminho para a pasta onde você deseja salvar a imagem.
		// Certifique-se de que a pasta existe e que você tem permissões de gravação.
		targetFolder := "/home/joao/uploadImagem/uploadImagem/img"
		img.Filename = "1.png"
		filePath := filepath.Join(targetFolder, img.Filename)

		// Salva a imagem na pasta especificada.
		err = img.Save(filePath)
		if err != nil {
			// Se ocorrer um erro durante a gravação da imagem, responde com o erro.
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		// Atualiza a variável global com a nova imagem.
		currentImage = img
		Envio(filePath)
		//	st, err := os.Stat(img)
		fmt.Println("linkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk:")
		fmt.Print(filePath)
		os.Remove(filePath)
		// Redireciona o usuário para a rota "/" após o upload bem-sucedido.
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Inicia o servidor na porta 8080, ouvindo e atendendo às requisições HTTP.
	r.Run(":8080")
}

func Envio(local string) {
	var envio bool

	emailUsuario := "jls.silva@discente.ufma.br"
	fmt.Printf(emailUsuario)

	from := "gerenciart.emserh@gmail.com"

	to := emailUsuario

	subject := ("Excelentíssima Fernanda novo documento")

	body := ("E Aí, Deu certo?")

	pdfPath := (local)

	pdfContent, err := ioutil.ReadFile(pdfPath)
	if err != nil {
		log.Printf("Erro ao ler arquivo PDF do funcionario")
		envio = false
	} else {
		envio = true
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "gerenciart.emserh@gmail.com"
	smtpPassword := "hikk wbqe dxap bwqc"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	var emailBuffer bytes.Buffer
	emailWriter := multipart.NewWriter(&emailBuffer)

	emailBuffer.WriteString("To: " + to + "\r\n")
	emailBuffer.WriteString("Subject: " + subject + "\r\n")
	emailBuffer.WriteString("MIME-Version: 1.0\r\n")
	emailBuffer.WriteString("Content-Type: multipart/mixed; boundary=" + emailWriter.Boundary() + "\r\n")
	emailBuffer.WriteString("\r\n")

	emailBuffer.WriteString("--" + emailWriter.Boundary() + "\r\n")
	emailBuffer.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	emailBuffer.WriteString("\r\n")
	emailBuffer.WriteString(body + "\r\n")

	emailBuffer.WriteString("--" + emailWriter.Boundary() + "\r\n")
	emailBuffer.WriteString("Content-Type: application/pdf; name=\"" + filepath.Base(pdfPath) + "\"\r\n")
	emailBuffer.WriteString("Content-Disposition: attachment; filename=\"" + filepath.Base(pdfPath) + "\"\r\n")
	emailBuffer.WriteString("Content-Transfer-Encoding: base64\r\n")
	emailBuffer.WriteString("\r\n")
	encoder := base64.NewEncoder(base64.StdEncoding, &emailBuffer)
	encoder.Write(pdfContent)
	encoder.Close()
	emailBuffer.WriteString("\r\n")

	emailBuffer.WriteString("--" + emailWriter.Boundary() + "--\r\n")

	if envio == true {
		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, emailBuffer.Bytes())
	}

	if err != nil {
		log.Printf("Erro ao enviar e-mail do funcionario: %s\n%s\n %s \n\n")
	}
}
