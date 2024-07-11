package main

import (
	"os"
)

var cpu = CPU{}
var ppu = PPU{}

func main() {
	if os.Args[1] == "--input" {
		if os.Args[2] == "asm" {
			cpu.read_asm()
		} else if os.Args[2] == "rom" {
			cpu.LoadNESFile(os.Args[3], ppu)
		}
	}
}
