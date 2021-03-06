package main

import (
	"a4"
	"bufio"
	"corefont"
	"fmt"
	"imagetype"
	"log"
	"os"
	"pdfjet"
	"pdfjet/src/compliance"
	"strings"
	"time"
)

// Example36 shows how you can add pages to PDF in random order.
func Example36() {
	file, err := os.Create("Example_36.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	pdf := pdfjet.NewPDF(w, compliance.PDF15)

	file1, err := os.Open("images/ee-map.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	image1 := pdfjet.NewImage(pdf, reader, imagetype.PNG)

	file2, err := os.Open("images/fruit.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	reader = bufio.NewReader(file2)
	image2 := pdfjet.NewImage(pdf, reader, imagetype.JPG)

	file3, err := os.Open("images/mt-map.bmp")
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()
	reader = bufio.NewReader(file3)
	image3 := pdfjet.NewImage(pdf, reader, imagetype.BMP)

	f1 := pdfjet.NewCoreFont(pdf, corefont.Helvetica())

	page1 := pdfjet.NewPage(pdf, a4.Portrait, false)

	text := pdfjet.NewTextLine(f1, "The map below is an embedded PNG image")
	text.SetLocation(90.0, 30.0)
	xy1 := text.DrawOn(page1)

	image1.SetLocation(90.0, xy1[1]+10.0)
	image1.ScaleBy(2.0 / 3.0)
	xy2 := image1.DrawOn(page1)

	text.SetText("JPG image file embedded once and drawn 3 times")
	text.SetLocation(90.0, xy2[1]+10.0)
	xy3 := text.DrawOn(page1)

	image2.SetLocation(90.0, xy3[1]+10.0)
	image2.ScaleBy(0.5)
	xy4 := image2.DrawOn(page1)

	image2.SetLocation(xy4[0]+10.0, xy3[1]+10.0)
	image2.ScaleBy(0.5)
	image2.SetRotateCW90(true)
	xy5 := image2.DrawOn(page1)

	image2.SetLocation(xy5[0]+10.0, xy3[1]+10.0)
	image2.SetRotateCW90(false)
	image2.ScaleBy(0.5)
	xy6 := image2.DrawOn(page1)

	image3.SetLocation(xy6[0]+10.0, xy6[1]+10.0)
	image3.ScaleBy(0.5)
	image3.DrawOn(page1)

	page2 := pdfjet.NewPage(pdf, a4.Portrait, false)

	text.SetText("This page was created after the second one but it was drawn first!")
	text.SetLocation(90.0, 30.0)
	xy7 := text.DrawOn(page2)

	image1.SetLocation(90.0, xy7[1]+10.0)
	image1.DrawOn(page2)

	pdf.AddPage(page2)
	pdf.AddPage(page1)

	pdf.Complete()
}

func main() {
	start := time.Now()
	Example36()
	elapsed := time.Since(start).String()
	fmt.Printf("Example_36 => %s\n", elapsed[:strings.Index(elapsed, ".")])
}
