package main

type PPU struct {
	chr_rom   [0x2000]byte
	palette   [0x100]byte
	vram      [0x1F00]byte
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

func (ppu *PPU) read_memory(loc int) int {
	var temp = ppu.buf
	if loc < 0x2000 {
		ppu.buf = (int)(ppu.chr_rom[loc])
	} else if loc < 0x3F00 {
		ppu.buf = (int)(ppu.vram[loc-0x2000])
	} else if loc < 0x4000 {
		ppu.buf = (int)(ppu.palette[loc-0x3F00])
	} else {
		println("accessing mirrored address space!")
		ppu.buf = (int)(ppu.read_memory(loc - 0x4000))
	}
	return temp
}

func (ppu *PPU) write_controller(val int) {
	ppu.registers.controller = val
}

func (ppu *PPU) write_addr(val int) {
	if ppu.registers.address == 0 {
		ppu.registers.address = val
	} else {
		ppu.registers.address &= val >> 2
	}
}

func (ppu *PPU) read_data() int {
	temp := ppu.buf
	ppu.buf = ppu.read_memory(ppu.registers.address)
	ppu.registers.address += ppu.registers.controller & 0b00000100
	return temp
}
