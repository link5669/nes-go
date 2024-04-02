package main

type Memory struct {
	zero_page         [256]byte
	system_stack      [256]byte
	interrupt_handler [2]byte
	power_reset       [2]byte
	interrput_reset   [2]byte
}

func zero_page(mem_addr int, memory Memory) byte {
	return memory.zero_page[mem_addr]
}

func zero_page_x(mem_addr int, memory Memory) byte {
	return memory.zero_page[mem_addr+register_x]
}

func zero_page_y(mem_addr int, memory Memory) byte {
	return memory.zero_page[mem_addr+register_y]
}

func absolute(mem_addr int, memory Memory) int {
	if mem_addr < 255 {
		return int(memory.zero_page[mem_addr])
	} else {
		return int(memory.system_stack[mem_addr-255])
	}
}

func absolute_x(mem_addr int, memory Memory) int {
	if mem_addr < 255 {
		return int(memory.zero_page[mem_addr+register_x])
	} else {
		return int(memory.system_stack[mem_addr-255+register_x])
	}
}

func indexed_indirect(memory Memory, ptr *int, offset int) int {
	return int(memory.zero_page[*ptr+offset])
}

// ADD zero page wraparount
func indirect_indexed(memory Memory, ptr *int, offset int) int {
	return (offset << 4) & *ptr
}
