package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Text(40, 10, "Hello, world")
	pdf.Image("./back.png", 56, 40, 100, 0, false, "", 0, "")

	err := pdf.OutputFileAndClose("./file.pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}
}
