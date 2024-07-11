package main

import (
	"os"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

var cpu = CPU{}
var ppu = PPU{horizontal_mirroring: false}

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Go NES!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	for !win.Closed() {
		win.Update()
	}
	if os.Args[1] == "--input" {
		if os.Args[2] == "asm" {
			cpu.read_asm()
		} else if os.Args[2] == "rom" {
			cpu.LoadNESFile(os.Args[3], &ppu, win)
		}
	}
}
