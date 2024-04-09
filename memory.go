package main

import (
	"math/rand/v2"
	"strconv"
)

type Memory struct {
	zero_page         [0x0100 - 0x000]int
	system_stack      [0x0200 - 0x0100]int
	gen_memory        [0xFFFA - 0x0200]int
	interrupt_handler [0x2]int
	power_reset       [0x2]int
	interrupt_reset   [0x2]int
}

func (m *Memory) immediate_addr(mem_addr int) *int {
	return &m.zero_page[mem_addr]
}

func (m *Memory) zero_page_addr(mem_addr int) *int {
	if mem_addr == 0xFE {
		var rand_num = rand.IntN(100)
		return &rand_num
	}
	return &m.zero_page[mem_addr]
}

func (m *Memory) zero_page_x_addr(mem_addr int) *int {
	if mem_addr+register_x > 0xff {
		cycles += 2
		return m.zero_page_x_addr(mem_addr - 0x100)
	}
	return &m.zero_page[mem_addr+register_x]
}

func (m *Memory) zero_page_y_addr(mem_addr int) *int {
	if mem_addr+register_y > 0xff {
		cycles += 2
		return m.zero_page_y_addr(mem_addr - 0x100)
	}
	return &m.zero_page[mem_addr+register_y]
}

func (m *Memory) absolute_addr(mem_addr int) *int {
	if mem_addr < 0x100 {
		return &m.zero_page[mem_addr]
	} else if mem_addr < 512 {
		return &m.system_stack[mem_addr-0x100]
	} else {
		return &m.gen_memory[mem_addr-0x200]
	}
}

func (m *Memory) absolute_x_addr(mem_addr int) *int {
	if mem_addr+register_x < 0x100 {
		return &m.zero_page[mem_addr+register_x]
	} else if mem_addr+register_x < 0x200 {
		return &m.system_stack[mem_addr-0x100+register_x]
	} else {
		return &m.gen_memory[mem_addr-0x200+register_x]
	}
}

func (m *Memory) absolute_y_addr(mem_addr int) *int {
	if mem_addr+register_y < 0x100 {
		return &m.zero_page[mem_addr+register_y]
	} else if mem_addr+register_y < 0x200 {
		return &m.system_stack[mem_addr-0x100+register_y]
	} else {
		return &m.gen_memory[mem_addr-0x200+register_y]
	}
}

func (m *Memory) indirect(first_half string, second_half string) {
	flipped_str := second_half + first_half
	retval := string_to_int(flipped_str)
	println(retval, "implement indirect??")
}

func (m *Memory) indexed_indirect_addr(val int) *int {
	if val+register_x > 0xff {
		cycles += 2
		val -= 0x100
	}
	lower := strconv.FormatInt(int64(m.zero_page[val+register_x]), 16)
	var upper string
	if val == 0xff {
		upper = strconv.FormatInt(int64(m.zero_page[register_x]), 16)
		if m.zero_page[register_x] < 0x10 {
			upper = "0" + upper
		}
	} else {
		upper = strconv.FormatInt(int64(m.zero_page[val+register_x+1]), 16)
		if m.zero_page[val+register_x+1] < 0x10 {
			upper = "0" + upper
		}
	}
	//hex <-> dec 0xa 0x10??
	if m.zero_page[val+register_x] < 0x10 {
		lower = "0" + lower
	}

	ret_val := string_to_int("$" + upper + lower)
	return m.absolute_addr(ret_val)
}

func (m *Memory) indirect_indexed_addr(val int) *int {
	lower := strconv.FormatInt(int64(m.zero_page[val]), 16)
	upper := strconv.FormatInt(int64(m.zero_page[val+1]), 16)
	if m.zero_page[val] < 0x10 {
		lower = "0" + lower
	}
	if m.zero_page[val+1] < 0x10 {
		upper = "0" + upper
	}
	ret_val := string_to_int("$"+upper+lower) + register_y
	return m.absolute_addr(ret_val)
}
