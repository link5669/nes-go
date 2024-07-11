package main

import (
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
	return &m.zero_page[mem_addr]
}

func (m *Memory) zero_page_x_addr(mem_addr int, cpu *CPU) *int {
	if mem_addr+cpu.register_x > 0xff {
		return m.zero_page_x_addr(mem_addr-0x100, cpu)
	}
	return &m.zero_page[mem_addr+cpu.register_x]
}

func (m *Memory) zero_page_y_addr(mem_addr int, cpu *CPU) *int {
	if mem_addr+cpu.register_y > 0xff {
		return m.zero_page_y_addr(mem_addr-0x100, cpu)
	}
	return &m.zero_page[mem_addr+cpu.register_y]
}

func (m *Memory) absolute_addr(mem_addr int, cpu *CPU) *int {
	if mem_addr < 0x100 {
		return &m.zero_page[mem_addr]
	} else if mem_addr < 512 {
		return &m.system_stack[mem_addr-0x100]
	} else if mem_addr < 0x10000 {
		return &m.gen_memory[mem_addr-0x200]
	} else {
		return m.absolute_addr(mem_addr-0x10000, cpu)
	}
}

func (m *Memory) absolute_x_addr(mem_addr int, check bool, cpu *CPU) *int {
	if check {
		check_cycle(mem_addr, cpu.register_x, cpu)
	}
	if mem_addr+cpu.register_x < 0x100 {
		return &m.zero_page[mem_addr+cpu.register_x]
	} else if mem_addr+cpu.register_x < 0x200 {
		return &m.system_stack[mem_addr-0x100+cpu.register_x]
	} else if mem_addr+cpu.register_x < 0x10000 {
		return &m.gen_memory[mem_addr-0x200+cpu.register_x]
	} else {
		cpu.cycles++
		return m.absolute_x_addr(mem_addr-0x10000, false, cpu)
	}
}

func check_cycle(mem_addr int, offset int, cpu *CPU) {
	if (mem_addr&0b11111111)+offset > 0xFF {
		cpu.cycles++
	}
}

func (m *Memory) absolute_y_addr(mem_addr int, cpu *CPU) *int {
	if mem_addr+cpu.register_y < 0x100 {
		check_cycle(mem_addr, cpu.register_y, cpu)
		return &m.zero_page[mem_addr+cpu.register_y]
	} else if mem_addr+cpu.register_y < 0x200 {
		check_cycle(mem_addr, cpu.register_y, cpu)
		return &m.system_stack[mem_addr-0x100+cpu.register_y]
	} else if mem_addr+cpu.register_y < 0x10000 {
		check_cycle(mem_addr, cpu.register_y, cpu)
		return &m.gen_memory[mem_addr-0x200+cpu.register_y]
	} else {
		return m.absolute_y_addr(mem_addr-0x10000, cpu)
	}
}

func (m *Memory) indirect(first_half string, second_half string) {
	flipped_str := second_half + first_half
	retval := string_to_int(flipped_str)
	println(retval, "implement indirect??")
}

func (m *Memory) indexed_indirect_addr(val int, cpu *CPU) *int {
	if val+cpu.register_x > 0xff {
		val -= 0x100
	}
	lower := strconv.FormatInt(int64(m.zero_page[val+cpu.register_x]), 16)
	var upper string
	if val == 0xff {
		upper = strconv.FormatInt(int64(m.zero_page[cpu.register_x]), 16)
		if m.zero_page[cpu.register_x] < 0x10 {
			upper = "0" + upper
		}
	} else {
		upper = strconv.FormatInt(int64(m.zero_page[val+cpu.register_x+1]), 16)
		if m.zero_page[val+cpu.register_x+1] < 0x10 {
			upper = "0" + upper
		}
	}
	if m.zero_page[val+cpu.register_x] < 0x10 {
		lower = "0" + lower
	}

	ret_val := string_to_int("$" + upper + lower)
	return m.absolute_addr(ret_val, cpu)
}

func (m *Memory) indirect_indexed_addr(val int, cpu *CPU) *int {
	upper_val := val + 1
	if upper_val > 0xFF {
		upper_val -= 0x100
		cpu.cycles++
	}
	lower := strconv.FormatInt(int64(m.zero_page[val]), 16)
	upper := strconv.FormatInt(int64(m.zero_page[upper_val]), 16)
	if m.zero_page[val] < 0x10 {
		lower = "0" + lower
	}
	if m.zero_page[upper_val] < 0x10 {
		upper = "0" + upper
	}
	ret_val := string_to_int("$"+upper+lower) + cpu.register_y
	if ret_val > 0xFFFF {
		ret_val -= 0x10000
		cpu.cycles++
	}
	return m.absolute_addr(ret_val, cpu)
}
