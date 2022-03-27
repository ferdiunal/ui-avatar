package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ferdiunal/avatars/src"
	"github.com/gofiber/fiber/v2"
)

var (
	fontSize float64 = 0.5
	color    string  = "000000"
	bgColor  string  = "ffffff"
	format   string  = "png"
	maxSize  uint64  = 1024
	size     uint64  = 512
	length   uint64  = 2
	bold     int     = 400
)

const (
	hexColorRegexString = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
)

var (
	hexColorRegex = regexp.MustCompile(hexColorRegexString)
)

func isHEXColor(color string) bool {
	return hexColorRegex.MatchString(color)
}

func handler(c *fiber.Ctx) error {

	c.Set("Pragma", "public")
	c.Set("Acess-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	c.Set("Access-Control-Max-Age", "1814400")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
	c.Set("Cache-Control", "public, max-age=1814400")

	getSize, _ := strconv.ParseUint(c.Query("size", string(rune(size))), 10, 64)
	if getSize > maxSize {
		getSize = maxSize
	} else if getSize < 25 || getSize == 0 {
		getSize = size
	}
	getFontSize, _ := strconv.ParseFloat(c.Query("font-size", fmt.Sprintf("%f", fontSize)), 32)
	if getFontSize == 0 || getFontSize < 0.1 {
		getFontSize = fontSize
	} else if getFontSize > 1 {
		getFontSize = 1
	}

	getLength, _ := strconv.ParseUint(c.Query("length", string(rune(length))), 10, 64)

	if getLength == 0 {
		getLength = length
	}

	getBackground := fmt.Sprintf("#%s", c.Query("background", ""))
	if !isHEXColor(getBackground) {
		getBackground = fmt.Sprintf("#%s", bgColor)
	}
	getColor := fmt.Sprintf("#%s", c.Query("color", ""))
	if !isHEXColor(getColor) {
		getColor = fmt.Sprintf("#%s", color)
	}

	names := strings.Split(strings.ToUpper(c.Query("name", "John Doe")), " ")
	getName := ""

	for i := range names {
		name := []rune(names[i])
		if i <= int(getLength) && len(name) > 0 {
			getName += string(name[0])
		}
	}

	isRounded := c.Query("rounded", "false") == "true"
	isBold := c.Query("bold", "false") == "true"
	if isBold {
		bold = 600
	} else {
		bold = 400
	}

	image := src.Image{
		Bold:       bold,
		Rounded:    isRounded,
		Background: getBackground,
		Color:      getColor,
		FontSize:   float32(math.Round(float64(getSize) * getFontSize)),
		Format:     c.Query("format", format),
		Name:       getName,
		Upper:      c.Query("upper", "false") == "true",
		Size:       getSize,
	}

	if image.Format == "svg" {
		c.Set("Content-Type", "image/svg+xml")

		c.WriteString(image.GenerateSvg())

		return nil
	}

	img := image.GenerateImage()

	c.Set("Content-Type", "image/png")
	c.Set("Content-Length", strconv.Itoa(len(img.Bytes())))

	c.Write(img.Bytes())
	return nil
}

func main() {

	app := fiber.New()

	app.Get("/", handler)

	log.Fatal(app.Listen(":9000"))
}
