// https://skilldrick.github.io/easy6502/ -- use for testing!
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var program_counter int = 0
var stack_pointer int = 0
var accumulator int = 0
var register_x int = 0
var register_y int = 0
var cpu_status Status
var memory Memory

var head Program_Code = Program_Code{}

func accumulator_add(val int) {
	if accumulator+val > 255 {
		accumulator += val - 256
	} else {
		accumulator += val
	}

}

func string_to_int(str string) int {
	split := strings.TrimPrefix(str, "#")
	if len(strings.Split(split, "$")) > 1 {
		split = strings.TrimPrefix(split, "$")
		split = strings.TrimLeft(split, "0")
		value, err := strconv.ParseInt(split, 16, 64)
		if err != nil {
			fmt.Printf("Conversion failed: %s\n", err)
		} else {
			return int(value)
		}
	}
	val, er := strconv.Atoi(split)
	if er != nil {
		fmt.Printf("Conversion failed: %s\n", er)
	}
	return int(val)
}

func main() {
	f, err := os.Open("test.asm")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var ptr *Program_Code
	ptr = &head
	index := 0
	for scanner.Scan() {
		curr_line := scanner.Text()
		curr_line = strings.TrimLeft(curr_line, " ")
		ptr.index = index
		if strings.Contains(curr_line, ":") {
			ptr.code_type = function_definition
			ptr.destination = strings.Split(curr_line, ":")[0]
		} else {
			ptr.code_type = op_code
			split_line := strings.Split(curr_line, " ")
			if len(split_line) > 1 {
				if strings.Contains(split_line[1], "$") {
					split_addr_1 := strings.Split(split_line[1], ",")
					var input_addr_1 int
					if strings.Contains(split_line[1], "#") {
						ptr.addr_mode = immediate_type
						input_addr_1 = string_to_int(split_line[1])
					} else if len(split_addr_1) > 1 {
						input_addr_1 = string_to_int(split_addr_1[0])
						if split_addr_1[1] == "Y" {
							if len(strings.Split(split_addr_1[0], "")) == 4 {
								ptr.addr_mode = absolute_y_type
							} else {
								ptr.addr_mode = zero_page_y_type
							}
						} else if split_addr_1[1] == "X" {
							if len(strings.Split(split_addr_1[0], "")) == 4 {
								ptr.addr_mode = absolute_x_type
							} else {
								ptr.addr_mode = zero_page_x_type
							}
						}
					} else {
						input_addr_1 = string_to_int(split_line[1])
						if len(strings.Split(split_addr_1[0], "")) == 5 {
							ptr.addr_mode = absolute_type
						} else {
							ptr.addr_mode = zero_page_type
						}
					}
					switch split_line[0] {
					case "LDX":
						ptr.op_code = LDX_cmd
						ptr.mem_1 = input_addr_1
					case "LDA":
						ptr.op_code = LDA_cmd
						ptr.mem_1 = input_addr_1
					case "STA":
						ptr.op_code = STA_cmd
						ptr.mem_1 = input_addr_1
					case "ADC":
						ptr.op_code = ADC_cmd
						ptr.mem_1 = input_addr_1
					case "STY":
						ptr.op_code = STY_cmd
						ptr.mem_1 = input_addr_1
					case "STX":
						ptr.op_code = STX_cmd
						ptr.mem_1 = input_addr_1
					case "CPX":
						ptr.op_code = CPX_cmd
						ptr.mem_1 = input_addr_1
					case "CMP":
						ptr.op_code = CMP_cmd
						ptr.mem_1 = input_addr_1
					}
				} else {
					switch split_line[0] {
					case "BNE":
						ptr.code_type = op_code
						ptr.op_code = BNE_cmd
						ptr.destination = split_line[1]
					}
				}
			} else {
				split_line := strings.Split(curr_line, " ")
				switch split_line[0] {
				case "BRK":
					ptr.op_code = BRK_cmd
				case "TAX":
					ptr.op_code = TAX_cmd
				case "TAY":
					ptr.op_code = TAY_cmd
				case "TYA":
					ptr.op_code = TYA_cmd
				case "INX":
					ptr.op_code = INX_cmd
				case "INY":
					ptr.op_code = INY_cmd
				case "DEX":
					ptr.op_code = DEX_cmd
				case "DEY":
					ptr.op_code = DEY_cmd
				case "SEI":
					ptr.op_code = SEI_cmd
				case "SED":
					ptr.op_code = SED_cmd
				case "SEC":
					ptr.op_code = SEC_cmd
				case "CLC":
					ptr.op_code = CLC_cmd
				case "CLD":
					ptr.op_code = CLD_cmd
				case "CLI":
					ptr.op_code = CLI_cmd
				case "CLV":
					ptr.op_code = CLV_cmd
				}
			}
		}
		ptr.next = &Program_Code{}
		ptr = ptr.next
		index++
	}

	dump_program()
	run_program()
	dump_contents()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func run_program() {
	var ptr *Program_Code = &head
	for ptr.next != nil {
		if ptr.code_type == function_definition {
			//TODO
		} else {
			switch ptr.op_code {
			case LDA_cmd:
				if ptr.addr_mode == immediate_type {
					LDA(ptr.mem_1)
				}
			case LDX_cmd:
				if ptr.addr_mode == immediate_type {
					LDX(ptr.mem_1)
				}
			case CPX_cmd:
				if ptr.addr_mode == immediate_type {
					CPX(ptr.mem_1)
				}
			case CMP_cmd:
				if ptr.addr_mode == immediate_type {
					CMP(ptr.mem_1)
				}
			case STA_cmd:
				if ptr.addr_mode == zero_page_type {
					STA(memory.zero_page_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_x_type {
					STA(memory.zero_page_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_y_type {
					STA(memory.zero_page_y_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_type {
					STA(memory.absolute_addr(ptr.mem_1))
				}
			case BRK_cmd:
				// dump_contents()
				// os.Exit(0)
			case TAX_cmd:
				TAX()
			case TAY_cmd:
				TAY()
			case TYA_cmd:
				TYA()
			case INX_cmd:
				INX()
			case INY_cmd:
				INY()
			case DEX_cmd:
				DEX()
			case DEY_cmd:
				DEY()
			case SEI_cmd:
				SEI()
			case SED_cmd:
				SED()
			case SEC_cmd:
				SEC()
			case CLC_cmd:
				CLC()
			case CLD_cmd:
				CLD()
			case CLI_cmd:
				CLI()
			case CLV_cmd:
				CLV()
			case STX_cmd:
				if ptr.addr_mode == zero_page_type {
					STX(memory.zero_page_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_x_type {
					STX(memory.zero_page_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_type {
					STX(memory.absolute_addr(ptr.mem_1))
				}
			case STY_cmd:
				if ptr.addr_mode == zero_page_type {
					STY(memory.zero_page_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_x_type {
					STY(memory.zero_page_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_type {
					STY(memory.absolute_addr(ptr.mem_1))
				}
			case ADC_cmd:
				if ptr.addr_mode == immediate_type {
					ADC(ptr.mem_1)
				} else if ptr.addr_mode == zero_page_type {
					ADC(*memory.zero_page_addr(ptr.mem_1))
				}
			case BNE_cmd:
				if !cpu_status.zero_flag {
					var local_ptr = &head
					println(register_x)
					for local_ptr.next != nil {
						if local_ptr.destination == ptr.destination && local_ptr.code_type == function_definition {
							ptr = local_ptr
							break
						}
						local_ptr = local_ptr.next
					}
				}
			}
		}
		ptr = ptr.next
	}
}

func load(val int, desti *int) {
	*desti = val
	cpu_status.zero_flag = *desti == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 00
}

func LDA(mem_addr int) {
	load(mem_addr, &accumulator)
}

func LDX(mem_addr int) {
	load(mem_addr, &register_x)
}

func LDY(mem_addr int) {
	load(mem_addr, &register_y)
}

func STA(val *int) {
	*val = accumulator
}

func STX(mem_addr *int) {
	*mem_addr = register_x
}

func STY(mem_addr *int) {
	*mem_addr = register_y
}

func TAX() {
	register_x = accumulator
	cpu_status.zero_flag = register_x == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

func TAY() {
	register_y = accumulator
	cpu_status.zero_flag = register_y == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

func TXA() {
	accumulator = register_x
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func TYA() {
	accumulator = register_y
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

// func TSX() {
// 	register_x = memory[stack_pointer]
// 	cpu_status.zero_flag = register_x == 0
// 	cpu_status.negative_flag = register_x&0b10000000 != 0
// }

// func TXS() {
// 	memory[stack_pointer] = register_x
// }

func CPX(val int) {
	cpu_status.carry_flag = register_x >= val
	cpu_status.zero_flag = register_x == val
	cpu_status.negative_flag = (register_x-val)&0b1000000 != 0
}

func CMP(val int) {
	cpu_status.carry_flag = accumulator >= val
	cpu_status.zero_flag = accumulator == val
	cpu_status.negative_flag = (accumulator-val)&0b1000000 != 0
}

// func PHP() {
// 	store_status := Status{ // b == Student{"Bob", 0}
// 		carry_flag:        cpu_status.carry_flag,
// 		zero_flag:         cpu_status.zero_flag,
// 		interrupt_disable: cpu_status.interrupt_disable,
// 		decimal_mode:      cpu_status.decimal_mode,
// 		break_command:     true,
// 		overflow_flag:     cpu_status.overflow_flag,
// 		negative_flag:     cpu_status.negative_flag,
// 	}
// 	memory[stack_pointer] = cpu_status
// }

// func PLA() {
// 	accumulator = memory[stack_pointer]
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0
// }

// func PLP() {
// 	cpu_status.carry_flag = memory[stack_pointer].carry_flag
// 	cpu_status.zero_flag = memory[stack_pointer].zero_flag
// 	cpu_status.interrupt_disable = memory[stack_pointer].interrupt_disable
// 	cpu_status.decimal_mode = memory[stack_pointer].decimal_mode
// 	cpu_status.break_command = memory[stack_pointer].break_command
// 	cpu_status.overflow_flag = memory[stack_pointer].overflow_flag
// 	cpu_status.negative_flag = memory[stack_pointer].negative_flag
// }

// func AND(mem_addr int) {
// 	accumulator &= memory[mem_addr]
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0
// }

// func EOR(mem_addr int) {
// 	accumulator = (accumulator || memory[mem_addr]) && accumulator && memory[mem_addr]
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0
// }

// func ORA(mem_addr int) {
// 	accumulator |= memory[mem_addr]
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0
// }

// func BIT() {
// 	//TODO
// }

// func INC(mem_addr int) {
// 	memory[mem_addr]++
// 	cpu_status.zero_flag = memory[mem_addr] == 0
// 	cpu_status.negative_flag = memory[mem_addr]&0b10000000 != 0
// }

func INX() {
	register_x++
	cpu_status.zero_flag = register_x == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

func INY() {
	register_y++
	cpu_status.zero_flag = register_y == 0
	cpu_status.negative_flag = register_y&0b10000000 != 0
}

// func DEC(mem_addr int) {
// 	memory[mem_addr]--
// 	cpu_status.zero_flag = memory[mem_addr] == 0
// 	cpu_status.negative_flag = memory[mem_addr]&0b10000000 != 0
// }

func DEX() {
	register_x--
	cpu_status.zero_flag = register_x == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

func DEY() {
	register_y--
	cpu_status.zero_flag = register_y == 0
	cpu_status.negative_flag = register_y&0b10000000 != 0
}

// // supposed to be mem address or accum?
// func ASL() {
// 	cpu_status.carry_flag = accumulator&0b10000000 != 0
// 	accumulator &= 0b11111110
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0
// }

// func LSR() {
// 	cpu_status.carry_flag = accumulator&0b00000001 != 0
// 	accumulator /= 2
// 	accumulator |= 0b10000000
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0
// }

// func ROR() {
// 	accumulator = accumulator >> 1
// 	if cpu_status.carry_flag {
// 		accumulator |= 0b00000001
// 	} else {
// 		accumulator &= 0b11111110
// 	}
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b00000001 != 0 != 0
// }

// func ROL() {
// 	accumulator = accumulator << 1
// 	if cpu_status.carry_flag {
// 		accumulator |= 0b10000000
// 	} else {
// 		accumulator &= 0b01111111
// 	}
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0 != 0
// }

func CLC() {
	cpu_status.carry_flag = false
}

func CLD() {
	cpu_status.decimal_mode = false
}

func CLI() {
	cpu_status.interrupt_disable = false
}

func CLV() {
	cpu_status.overflow_flag = false
}

func SEC() {
	cpu_status.carry_flag = true
}

func SED() {
	cpu_status.decimal_mode = true
}

func SEI() {
	cpu_status.interrupt_disable = true
}

func ADC(val int) {
	if cpu_status.carry_flag {
		accumulator_add(val + 1)
	} else {
		accumulator_add(val)
	}
	cpu_status.carry_flag = accumulator&0b100000000 != 0
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func SBC(val int) {
	if !cpu_status.carry_flag {
		accumulator_add(-(val + 1))
	} else {
		accumulator_add(-val)
	}
	cpu_status.carry_flag = accumulator&0b100000000 != 0
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}
