package main

import "strconv"

type Memory struct {
	zero_page         [256]int
	system_stack      [256]int
	interrupt_handler [2]int
	power_reset       [2]int
	interrupt_reset   [2]int
}

func (m *Memory) immediate_addr(mem_addr int) *int {
	if mem_addr > 255 {
		return m.immediate_addr(mem_addr - 256)
	}
	return &m.zero_page[mem_addr]
}

func (m *Memory) zero_page_addr(mem_addr int) *int {
	if mem_addr > 255 {
		return m.zero_page_addr(mem_addr - 256)
	}
	return &m.zero_page[mem_addr]
}

func (m *Memory) zero_page_x_addr(mem_addr int) *int {
	if mem_addr+register_x > 255 {
		return m.zero_page_x_addr(mem_addr - 256)
	}
	return &m.zero_page[mem_addr+register_x]
}

func (m *Memory) zero_page_y_addr(mem_addr int) *int {
	if mem_addr+register_y > 255 {
		return m.zero_page_y_addr(mem_addr - 256)
	}
	return &m.zero_page[mem_addr+register_y]
}

func (m *Memory) absolute_addr(mem_addr int) *int {
	if mem_addr > 511 {
		return m.absolute_addr(mem_addr - 512)
	} else if mem_addr < 255 {
		return &m.zero_page[mem_addr]
	} else {
		return &m.system_stack[mem_addr-256]
	}
}

func (m *Memory) absolute_x_addr(mem_addr int) *int {
	if mem_addr > 511 {
		return m.absolute_x_addr(mem_addr - 512)
	} else if mem_addr < 255 {
		return &m.zero_page[mem_addr+register_x]
	} else {
		return &m.system_stack[mem_addr-256+register_x]
	}
}

func (m *Memory) absolute_y_addr(mem_addr int) *int {
	if mem_addr > 511 {
		return m.absolute_y_addr(mem_addr - 512)
	} else if mem_addr < 255 {
		return &m.zero_page[mem_addr+register_y]
	} else {
		return &m.system_stack[mem_addr-256+register_y]
	}
}

func (m *Memory) indirect(first_half string, second_half string) {
	flipped_str := second_half + first_half
	retval := string_to_int(flipped_str)
	println(retval, "implement indirect??")
}

func (m *Memory) indexed_indirect_addr(val int) *int {
	lower := strconv.Itoa(m.zero_page[val+register_x])
	upper := strconv.Itoa(m.zero_page[val+register_x+1])
	if m.zero_page[val+register_x] < 10 {
		upper = "0" + upper
	}
	if m.zero_page[val+register_x+1] < 10 {
		lower = "0" + lower
	}
	ret_val := string_to_int("$" + upper + lower)
	return m.absolute_addr(ret_val)
}

func (m *Memory) indirect_indexed_addr(val int) *int {
	lower := strconv.Itoa(m.zero_page[val])
	upper := strconv.Itoa(m.zero_page[val])
	if m.zero_page[val+register_x] < 10 {
		upper = "0" + upper
	}
	if m.zero_page[val+register_x+1] < 10 {
		lower = "0" + lower
	}
	ret_val := string_to_int("$"+upper+lower) + register_y
	return m.absolute_addr(ret_val)
}
