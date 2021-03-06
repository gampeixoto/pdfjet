package main

import (
	"bufio"
	"corefont"
	"fmt"
	"imagetype"
	"letter"
	"log"
	"os"
	"pdfjet"
	"pdfjet/src/color"
	"pdfjet/src/compliance"
	"strings"
	"time"
)

// Example30 draws the Canadian flag using a Path object that contains both lines
// and curve segments. Every curve segment must have exactly 2 control points.
func Example30() {
	file, err := os.Create("Example_30.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	pdf := pdfjet.NewPDF(w, compliance.PDF15)

	font := pdfjet.NewCoreFont(pdf, corefont.Helvetica())

	file1, err := os.Open("images/map407.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	image1 := pdfjet.NewImage(pdf, reader, imagetype.PNG)
	image1.SetLocation(10.0, 100.0)

	file2, err := os.Open("images/qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	reader = bufio.NewReader(file2)
	image2 := pdfjet.NewImage(pdf, reader, imagetype.PNG)
	image2.SetLocation(10.0, 100.0)

	// Create the first page after all the resources have been added to the PDF.
	page := pdfjet.NewPage(pdf, letter.Portrait, true)

	textLine := pdfjet.NewTextLine(font, "© OpenStreetMap contributors")
	textLine.SetLocation(430.0, 655.0)
	xy := textLine.DrawOn(page)

	uri := "http://www.openstreetmap.org/copyright"
	textLine = pdfjet.NewTextLine(font, "http://www.openstreetmap.org/copyright")
	textLine.SetURIAction(&uri)
	textLine.SetLocation(380.0, xy[1]+font.GetHeight())
	textLine.DrawOn(page)

	group := pdfjet.NewOptionalContentGroup("Map")
	group.Add(image1)
	group.SetVisible(true)
	// group.SetPrintable(true)
	group.DrawOn(page)

	textBox := pdfjet.NewTextBox(font)
	textBox.SetText("Hello Blue Layer Text")
	textBox.SetLocation(300.0, 200.0)

	line := pdfjet.NewLine(300.0, 250.0, 500.0, 250.0)
	line.SetWidth(2.0)
	line.SetColor(color.Blue)

	group = pdfjet.NewOptionalContentGroup("Blue")
	group.Add(textBox)
	group.Add(line)
	// group.SetVisible(true)
	group.DrawOn(page)

	line = pdfjet.NewLine(300.0, 260.0, 500.0, 260.0)
	line.SetWidth(2.0)
	line.SetColor(color.Red)

	group = pdfjet.NewOptionalContentGroup("Barcode")
	group.Add(image2)
	group.Add(line)
	group.SetVisible(true)
	group.SetPrintable(true)
	group.DrawOn(page)

	pdf.Complete()
}

func main() {
	start := time.Now()
	Example30()
	elapsed := time.Since(start).String()
	fmt.Printf("Example_30 => %s\n", elapsed[:strings.Index(elapsed, ".")])
}
