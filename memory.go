package main

type Memory struct {
	zero_page         [256]int
	system_stack      [256]int
	interrupt_handler [2]int
	power_reset       [2]int
	interrupt_reset   [2]int
}

func (m *Memory) immediate_addr(mem_addr int) *int {
	return &m.zero_page[mem_addr]
}

func (m *Memory) zero_page_addr(mem_addr int) *int {
	return &(m.zero_page[mem_addr])
}

func (m *Memory) zero_page_x_addr(mem_addr int) *int {
	return &m.zero_page[mem_addr+register_x]
}

func (m *Memory) zero_page_y_addr(mem_addr int) *int {
	return &m.zero_page[mem_addr+register_y]
}

func (m *Memory) absolute_addr(mem_addr int) *int {
	if mem_addr < 255 {
		return &m.zero_page[mem_addr]
	} else {
		return &m.system_stack[mem_addr-255]
	}
}

func (m *Memory) absolute_x_addr(mem_addr int) *int {
	if mem_addr < 255 {
		return &m.zero_page[mem_addr+register_x]
	} else {
		return &m.system_stack[mem_addr-255+register_x]
	}
}

func (m *Memory) indexed_indirect_addr(ptr *int, offset int) *int {
	return &m.zero_page[*ptr+offset]
}

// ADD zero page wraparount
// func (m *Memory) indirect_indexed_addr(ptr *int, offset int) *int {
// 	return (offset << 4) & *ptr
// }
