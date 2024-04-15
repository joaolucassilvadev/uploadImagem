package main

func Service() { /*
		f := "jls.silva@discente.ufma.br"
		// Endereço de e-mail do remetente
		from := "js0454261@gmail.com"
		// Endereço de e-mail para onde enviar
		to := f
		// Assunto do e-mail
		subject := "Teste de e-mail com anexo"
		// Corpo do e-mail
		body := "Corpo do e-mail"

		cpf := d.Cpf
		pdfPath := ("/home/joao/login_portal/login_portal_colaborador/service/paginas/" + cpf + ".pdf")

		// Lendo o conteúdo do arquivo PDF
		pdfContent, err := ioutil.ReadFile(pdfPath)
		if err != nil {
			log.Fatalf("Erro ao ler arquivo PDF: %s", err)
		}

		// Configuração do SMTP
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"
		smtpUsername := "js0454261@gmail.com"
		smtpPassword := "lhts rsjs onnt wyme"

		// Autenticação SMTP
		auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

		// Criando buffer para o corpo do e-mail
		var emailBuffer bytes.Buffer
		emailWriter := multipart.NewWriter(&emailBuffer)

		// Escrevendo o cabeçalho do e-mail
		emailBuffer.WriteString("To: " + to + "\r\n")
		emailBuffer.WriteString("Subject: " + subject + "\r\n")
		emailBuffer.WriteString("MIME-Version: 1.0\r\n")
		emailBuffer.WriteString("Content-Type: multipart/mixed; boundary=" + emailWriter.Boundary() + "\r\n")
		emailBuffer.WriteString("\r\n")

		// Escrevendo o corpo do e-mail
		emailBuffer.WriteString("--" + emailWriter.Boundary() + "\r\n")
		emailBuffer.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
		emailBuffer.WriteString("\r\n")
		emailBuffer.WriteString(body + "\r\n")

		// Escrevendo o anexo do PDF
		emailBuffer.WriteString("--" + emailWriter.Boundary() + "\r\n")
		emailBuffer.WriteString("Content-Type: application/pdf; name=\"" + filepath.Base(pdfPath) + "\"\r\n")
		emailBuffer.WriteString("Content-Disposition: attachment; filename=\"" + filepath.Base(pdfPath) + "\"\r\n")
		emailBuffer.WriteString("Content-Transfer-Encoding: base64\r\n")
		emailBuffer.WriteString("\r\n")
		encoder := base64.NewEncoder(base64.StdEncoding, &emailBuffer)
		encoder.Write(pdfContent)
		encoder.Close()
		emailBuffer.WriteString("\r\n")

		// Escrevendo o final do e-mail
		emailBuffer.WriteString("--" + emailWriter.Boundary() + "--\r\n")

		// Envia o e-mail
		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, emailBuffer.Bytes())
		if err != nil {
			log.Fatalf("Erro ao enviar e-mail: %s", err)
		} */
}
