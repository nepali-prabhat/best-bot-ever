package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/golang/freetype/truetype"
	"github.com/lukesampson/figlet/figletlib"
	img_font "golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func handleShow(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	// Read the font data
	fontBytes, fontErr := ioutil.ReadFile("./JetBrainsMomo.ttf")
	if fontErr != nil {
		log.Println(fontErr)
		return
	}
	f, ttErr := truetype.Parse(fontBytes)
	if ttErr != nil {
		log.Println(ttErr)
		return
	}

	message := strings.TrimSpace(command[len("show"):])
	if font == nil {
		return
	}
	var tempMsg strings.Builder
	figletlib.FPrintMsg(&tempMsg, message, font, 80, font.Settings(), "center")
	msg := tempMsg.String()
	text := strings.Split(msg, "\n")

	// some constants for image properties
	const (
		size      = 6
		dpi       = 300
		spacing   = 1
		charWidth = 15 // subject to change with size and spacing
	)

	// calculating height dynamically
	y := 0 + int(math.Ceil(size*dpi/72))
	dy := int(math.Ceil(size * spacing * dpi / 72))
	imgH := dy * len(text)
	const imgW = charWidth * 80

	fg, bg := image.White, image.Transparent

	rgba := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	// measurement
	// vert := charWidth
	// for i := 0; i < 200; i++ {
	// 	for j := 1; j < imgW/vert; j++ {
	// 		rgba.Set(vert*j, 10+i, ruler)
	// 	}
	// 	rgba.Set(10+i, 10, ruler)
	// }

	d := &img_font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: img_font.HintingNone,
		}),
	}
	for _, s := range text {
		d.Dot = fixed.P(0, y)
		d.DrawString(s)
		y += dy
	}
	// Save that RGBA image to disk.
	fileName := m.Author.Username + "_" + time.Now().String() + "_" + ".png"

	var buff bytes.Buffer

	b := bufio.NewWriter(&buff)
	pngEncodeErr := png.Encode(b, rgba)
	if pngEncodeErr != nil {
		fmt.Println("Error creating png file, ", pngEncodeErr)
		return
	}
	flushErr := b.Flush()
	if flushErr != nil {
		fmt.Println("png buffer flush err, ", flushErr)
	}
	reader := bytes.NewReader(buff.Bytes())
	_, err := s.ChannelFileSendWithMessage(m.ChannelID, "", fileName, reader)
	if err != nil {
		fmt.Println("Error saying, ", err)
	}
}
