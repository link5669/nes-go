package main

type PPU struct {
	chr_rom   [0xFFFF]byte
	palette   [0x20]byte
	vram      [0x80]byte
	oam_data  [0x10]byte
	buf       int
	registers PPU_Registers
}

type PPU_Registers struct {
	controller int
	mask       int
	status     int
	oam_addr   int
	oam_data   int
	scroll     int
	address    int
	data       int
	oam_dma    int
}

func (ppu *PPU) write_addr(val int) {
	if ppu.registers.address == 0 {
		ppu.registers.address = val
	} else {
		ppu.registers.address &= val >> 2
	}
}

func (ppu *PPU) read_data(val int, mem Memory) int {
	if ppu.buf == 0 {
		ppu.buf = *mem.absolute_addr(ppu.registers.address)
		ppu.registers.address += ppu.registers.controller & 0b00000100
		return 0
	} else {
		return ppu.buf
	}
}
