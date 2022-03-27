package src

import (
	"bytes"
	"fmt"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type Image struct {
	Rounded    bool
	Background string
	Color      string
	FontSize   float32
	Format     string
	Name       string
	Upper      bool
	Size       uint64
	Bold       int
}

func (i *Image) GenerateImage() *bytes.Buffer {

	half := float64(i.Size) / 2
	dc := gg.NewContext(int(i.Size), int(i.Size))
	dc.SetHexColor(i.Background)
	if i.Rounded {
		dc.DrawCircle(half, half, half)
	} else {
		dc.DrawRectangle(0, 0, float64(i.Size), float64(i.Size))
	}
	dc.Fill()

	dc.Clip()
	dc.InvertMask()

	dc.SetHexColor(i.Color)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(i.FontSize),
		DPI:  72,
	})
	dc.SetFontFace(face)

	_, h := dc.MeasureString(i.Name)
	dc.DrawStringWrapped(i.Name, half, half-h/2, 0.5, 0.15, float64(i.Size), 1.5, gg.AlignCenter)
	dc.Fill()

	buf := new(bytes.Buffer)
	dc.EncodePNG(buf)
	return buf
}

func (i *Image) GenerateSvg() string {
	width := fmt.Sprintf("%dpx", i.Size)
	height := fmt.Sprintf("%dpx", i.Size)
	half := i.Size / 2

	svg := fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"%s\" height=\"%s\" viewBox=\"0 0 %d %d\" version=\"1.1\">\n", width, height, i.Size, i.Size)
	shape := "rect"
	if i.Rounded {
		shape = "circle"
	}
	svg += fmt.Sprintf("<%s xmlns=\"http://www.w3.org/2000/svg\" fill=\"%s\" cx=\"%d\" width=\"%s\" height=\"%s\" cy=\"%d\" r=\"%d\" />\n", shape, i.Background, half, width, height, half, half)
	svg += fmt.Sprintf("<text xmlns=\"http://www.w3.org/2000/svg\" x=\"%s\" y=\"%s\" style=\"color: %s; line-height: 1;font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Arial', 'Ubuntu', 'Helvetica Neue', sans-serif;\" alignment-baseline=\"middle\" text-anchor=\"middle\" font-size=\"%f\" font-weight=\"%d\" dy=\".1em\" dominant-baseline=\"middle\" fill=\"%s\">%s</text>\n", "50%", "50%", i.Color, i.FontSize, i.Bold, i.Color, i.Name)
	svg += "</svg>"

	return svg
}
