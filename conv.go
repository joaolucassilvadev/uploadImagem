package main

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func conv() {
	// Crie um novo documento PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Adicione uma nova página ao documento
	pdf.AddPage()

	// Adicione suas imagens ao documento
	inputPaths := []string{"/home/joao/convImgtopdf/image-pdf/Captura de tela 2023-11-23 110159.png"} // Caminho adicionado

	for _, imgPath := range inputPaths {
		pdf.Image(imgPath, 0, 0, 210, 0, false, "", 0, "")
		pdf.AddPage() // Adicione uma nova página para cada imagem
	}

	// Salve o documento em um arquivo PDF
	outputPath := "/home/joao/convImgtopdf/image-pdf/imaaag.pdf"
	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo PDF:", err)
		os.Exit(1)
	}

	fmt.Printf("PDF gerado com sucesso. Arquivo: %s\n", outputPath)
}
