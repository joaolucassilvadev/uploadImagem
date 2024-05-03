package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/smtp"
	"path/filepath"
)

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
