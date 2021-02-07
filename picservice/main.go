package main

import (
	_ "image/png"
	"os"
	"time"

	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/texture"

	log "github.com/sirupsen/logrus"
)

func main() {
	// load a PNG specified on the command line (8x8 recommended).
	if len(os.Args) < 2 {
		log.Fatal("specify a png image to display.")
	}

	tx, err := texture.Load(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fb := screen.NewFrameBuffer()
	texture.Blit(fb.Texture, 0, 0, tx, 0, 0, 8, 8)
	screen.Draw(fb)

	time.Sleep(3000 * time.Millisecond)
	screen.Clear()
}
