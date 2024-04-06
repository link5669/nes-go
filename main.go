// https://skilldrick.github.io/easy6502/ -- use for testing!
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

var program_counter int = 0x600
var stack_pointer int = 0xFF
var accumulator int = 0
var register_x int = 0
var register_y int = 0
var cpu_status Status
var memory Memory

var const_head Const = Const{}

var head Program_Code = Program_Code{}

func accumulator_add(val int) {
	if accumulator+val > 0xFF {
		accumulator += val - 0x100
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

func split(val string, char string) []string {
	split_val := strings.Split(val, char)
	ret_val := []string{}
	for i := 0; i < len(split_val); i++ {
		if !(split_val[i] == "" || split_val[i] == " ") {
			ret_val = append(ret_val, split_val[i])
		}
	}
	return ret_val
}

func get_addr_mode_val(split_line []string) (addr_mode, int) {
	var ret_addr_mode addr_mode
	var ret_value int
	split_addr_1 := split(split_line[1], ",")
	if strings.Contains(split_line[1], "#") {
		ret_addr_mode = immediate_type
		ret_value = string_to_int(split_line[1])
	} else if len(split_addr_1) > 1 && !strings.Contains(split_line[1], "(") {
		ret_value = string_to_int(split_addr_1[0])
		if strings.ToUpper(split_addr_1[1]) == "Y" {
			if len(split(split_addr_1[0], "")) == 5 {
				ret_addr_mode = absolute_y_type
			} else {
				ret_addr_mode = zero_page_y_type
			}
		} else if strings.ToUpper(split_addr_1[1]) == "X" {
			if len(split(split_addr_1[0], "")) == 5 {
				ret_addr_mode = absolute_x_type
			} else {
				ret_addr_mode = zero_page_x_type
			}
		}
	} else {
		if !strings.Contains(split_line[1], "(") {
			ret_value = string_to_int(split_line[1])
			if len(split(split_addr_1[0], "")) == 5 {
				ret_addr_mode = absolute_type
			} else {
				ret_addr_mode = zero_page_type
			}
		} else {
			addr := strings.TrimLeft(split_line[1], "(")
			splita := split(addr, "")
			if splita[len(splita)-1] == ")" {
				ret_addr_mode = indexed_indirect_type
				addr = split(addr, ",")[0]
				ret_value = string_to_int(addr)
			} else {
				addr = split(addr, ",")[0]
				addr = strings.TrimRight(addr, ")")
				ret_value = string_to_int(addr)
				ret_addr_mode = indirect_indexed_type
			}
		}
	}
	return ret_addr_mode, ret_value
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
	var var_ptr *Const
	var_ptr = &const_head
	index := 0
	for scanner.Scan() {
		curr_line := scanner.Text()
		curr_line = strings.TrimSpace(curr_line)
		if curr_line == "" {
			continue
		} else if strings.Contains(curr_line, "define") {
			split_line := split(curr_line, " ")
			var_ptr.name = split_line[1]
			if split_line[1] == "sysRandom" {
				var_ptr.val = strconv.FormatInt(int64(rand.IntN(255)), 16)
				var_ptr.val = "#$" + var_ptr.val
			} else {
				var_ptr.val = split_line[2]
			}
			var_ptr.next = &Const{}
			var_ptr = var_ptr.next
			continue
		}
		ptr.index = index
		program_counter++
		if strings.Contains(curr_line, ":") {
			ptr.code_type = function_definition
			ptr.destination = split(curr_line, ":")[0]
		} else {
			ptr.code_type = op_code
			split_line := split(curr_line, " ")
			split_line[0] = strings.ToUpper(split_line[0])
			if len(split_line) > 1 {
				if !strings.Contains(split_line[1], "$") {
					var local_ptr = &const_head
					for local_ptr.next != nil {
						split_line[1] = strings.Replace(split_line[1], local_ptr.name, local_ptr.val, 1)
						local_ptr = local_ptr.next
					}
				}
				if strings.Contains(split_line[1], "$") || strings.Contains(split_line[1], "#") {
					var input_addr_1 int
					ptr.addr_mode, input_addr_1 = get_addr_mode_val(split_line)
					switch split_line[0] {
					case "LSR":
						ptr.op_code = LSR_cmd
						ptr.mem_1 = input_addr_1
					case "AND":
						ptr.op_code = AND_cmd
						ptr.mem_1 = input_addr_1
					case "DEC":
						ptr.op_code = DEC_cmd
						ptr.mem_1 = input_addr_1
					case "LDX":
						ptr.op_code = LDX_cmd
						ptr.mem_1 = input_addr_1
					case "LDY":
						ptr.op_code = LDY_cmd
						ptr.mem_1 = input_addr_1
					case "LDA":
						ptr.op_code = LDA_cmd
						ptr.mem_1 = input_addr_1
					case "BIT":
						ptr.op_code = BIT_cmd
						ptr.mem_1 = input_addr_1
					case "SBC":
						ptr.op_code = SBC_cmd
						ptr.mem_1 = input_addr_1
					case "STA":
						ptr.op_code = STA_cmd
						ptr.mem_1 = input_addr_1
					case "ADC":
						ptr.op_code = ADC_cmd
						ptr.mem_1 = input_addr_1
					case "INC":
						ptr.op_code = INC_cmd
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
					case "CPY":
						ptr.op_code = CPY_cmd
						ptr.mem_1 = input_addr_1
					case "JMP":
						//copied from below!
						ptr.op_code = JMP_cmd
						if strings.Contains(split_line[1], "(") {
							ptr.addr_mode = indirect_type
							addr := strings.TrimLeft(split_line[1], "(")
							addr = strings.TrimRight(addr, ")")
							ptr.destination = addr
						} else {
							ptr.addr_mode = absolute_type
							ptr.destination = split_line[1]
						}
					}
				} else {
					ptr.code_type = op_code
					switch split_line[0] {
					case "BNE":
						ptr.op_code = BNE_cmd
						ptr.destination = split_line[1]
					case "BEQ":
						ptr.op_code = BEQ_cmd
						ptr.destination = split_line[1]
					case "BCS":
						ptr.op_code = BCS_cmd
						ptr.destination = split_line[1]
					case "BCC":
						ptr.op_code = BCC_cmd
						ptr.destination = split_line[1]
					case "BPL":
						ptr.op_code = BPL_cmd
						ptr.destination = split_line[1]
					case "JSR":
						ptr.op_code = JSR_cmd
						ptr.destination = split_line[1]
					case "JMP":
						//copied from above!
						ptr.op_code = JMP_cmd
						if strings.Contains(split_line[1], "(") {
							ptr.addr_mode = indirect_type
							addr := strings.TrimLeft(split_line[1], "(")
							addr = strings.TrimRight(addr, ")")
							ptr.destination = addr
						} else {
							ptr.addr_mode = absolute_type
							ptr.destination = split_line[1]
						}
					}
				}
			} else {
				split_line := split(curr_line, " ")
				split_line[0] = strings.ToUpper(split_line[0])
				switch split_line[0] {
				case "PHP":
					ptr.op_code = PHP_cmd
				case "PLP":
					ptr.op_code = PLP_cmd
				case "LSR":
					ptr.op_code = LSR_cmd
				case "RTS":
					ptr.op_code = RTS_cmd
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
				case "PLA":
					ptr.op_code = PLA_cmd
				case "PHA":
					ptr.op_code = PHA_cmd
				case "TXS":
					ptr.op_code = TXS_cmd
				case "TSX":
					ptr.op_code = TSX_cmd
				case "TXA":
					ptr.op_code = TXA_cmd
				case "NOP":
					ptr.op_code = NOP_cmd
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
			//for real vals
			var val int = 0
			if ptr.addr_mode == immediate_type {
				val = ptr.mem_1
			} else if ptr.addr_mode == zero_page_type {
				val = *memory.zero_page_addr(ptr.mem_1)
			} else if ptr.addr_mode == zero_page_x_type {
				val = (*memory.zero_page_x_addr(ptr.mem_1))
			} else if ptr.addr_mode == zero_page_y_type {
				val = (*memory.zero_page_y_addr(ptr.mem_1))
			} else if ptr.addr_mode == absolute_type {
				val = (*memory.absolute_addr(ptr.mem_1))
			} else if ptr.addr_mode == absolute_x_type {
				val = (*memory.absolute_x_addr(ptr.mem_1))
			} else if ptr.addr_mode == absolute_y_type {
				val = (*memory.absolute_y_addr(ptr.mem_1))
			} else if ptr.addr_mode == indexed_indirect_type {
				val = (*memory.indexed_indirect_addr(ptr.mem_1))
			} else if ptr.addr_mode == indirect_indexed_type {
				val = (*memory.indirect_indexed_addr(ptr.mem_1))
			}
			switch ptr.op_code {
			case DEC_cmd:
				if ptr.addr_mode == zero_page_type {
					DEC(memory.zero_page_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_x_type {
					DEC(memory.zero_page_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_type {
					DEC(memory.absolute_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_x_type {
					DEC(memory.absolute_x_addr(ptr.mem_1))
				}
			case BIT_cmd:
				BIT(val)
			case INC_cmd:
				if ptr.addr_mode == zero_page_type {
					INC(memory.zero_page_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_x_type {
					INC(memory.zero_page_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_type {
					INC(memory.absolute_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_x_type {
					INC(memory.absolute_x_addr(ptr.mem_1))
				}
			case LDA_cmd:
				LDA(val)
			case SBC_cmd:
				SBC(val)
			case LDX_cmd:
				LDX(val)
			case LDY_cmd:
				LDY(val)
			case CPX_cmd:
				CPX(val)
			case CMP_cmd:
				CMP(val)
			case CPY_cmd:
				CPY(val)
			case STA_cmd:
				if ptr.addr_mode == zero_page_type {
					STA(memory.zero_page_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_x_type {
					STA(memory.zero_page_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_y_type {
					STA(memory.zero_page_y_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_type {
					STA(memory.absolute_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_x_type {
					STA(memory.absolute_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_y_type {
					STA(memory.absolute_y_addr(ptr.mem_1))
				} else if ptr.addr_mode == indexed_indirect_type {
					STA(memory.indexed_indirect_addr(ptr.mem_1))
				} else if ptr.addr_mode == indirect_indexed_type {
					STA(memory.indirect_indexed_addr(ptr.mem_1))
				}
			case AND_cmd:
				AND(val)
			case LSR_cmd:
				if ptr.addr_mode == zero_page_type {
					LSR(memory.zero_page_addr(ptr.mem_1))
				} else if ptr.addr_mode == zero_page_x_type {
					LSR(memory.zero_page_x_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_type {
					LSR(memory.absolute_addr(ptr.mem_1))
				} else if ptr.addr_mode == absolute_x_type {
					LSR(memory.absolute_x_addr(ptr.mem_1))
				} else {
					LSR(nil)
				}
			case PLP_cmd:
				PLP()
			case PHP_cmd:
				PHP()
			case BRK_cmd:
				dump_contents()
				os.Exit(0)
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
			case NOP_cmd:
				time.Sleep(1 * time.Millisecond)
			case PLA_cmd:
				PLA()
			case PHA_cmd:
				PHA()
			case TXS_cmd:
				TXS()
			case TSX_cmd:
				TSX()
			case TXA_cmd:
				TXA()
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
				ADC(val)
			case BNE_cmd:
				if !cpu_status.zero_flag {
					ptr = break_to(ptr.destination)
				}
			case BEQ_cmd:
				if cpu_status.zero_flag {
					ptr = break_to(ptr.destination)
				}
			case BCS_cmd:
				if cpu_status.carry_flag {
					ptr = break_to(ptr.destination)
				}
			case BCC_cmd:
				if !cpu_status.carry_flag {
					ptr = break_to(ptr.destination)
				}
			case BPL_cmd:
				if !cpu_status.negative_flag {
					ptr = break_to(ptr.destination)
				}
			case JMP_cmd:
				if ptr.addr_mode == indirect_type {
					int_addr := string_to_int(ptr.destination)
					val_1 := *memory.absolute_addr(int_addr)
					val_2 := *memory.absolute_addr(int_addr + 1)
					str_val_1 := strconv.FormatInt(int64(val_1), 16)
					str_val_2 := strconv.FormatInt(int64(val_2), 16)
					if val_1 < 10 {
						str_val_1 = "0" + str_val_1
					}
					full_addr := str_val_2 + str_val_1
					println("implement indirect jmp!", full_addr)
				} else if ptr.addr_mode == absolute_type {
					ptr = break_to(ptr.destination)
				}
			case JSR_cmd:
				stack_pointer--
				memory.system_stack[stack_pointer] = ptr.index
				var local_ptr = &head
				for local_ptr.next != nil {
					if local_ptr.destination == ptr.destination && local_ptr.code_type == function_definition {
						ptr = local_ptr
						break
					}
					local_ptr = local_ptr.next
				}
			case RTS_cmd:
				last_index := memory.system_stack[stack_pointer]
				stack_pointer++
				var local_ptr = &head
				for local_ptr.next != nil {
					if local_ptr.index == last_index {
						ptr = local_ptr
						break
					}
					local_ptr = local_ptr.next
				}
				print_screen()
			}
		}
		ptr = ptr.next
	}
}

func break_to(destination string) *Program_Code {
	var local_ptr = &head
	for local_ptr.next != nil {
		if local_ptr.destination == destination && local_ptr.code_type == function_definition {
			return local_ptr
		}
		local_ptr = local_ptr.next
	}
	return nil
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

func TSX() {
	register_x = memory.system_stack[stack_pointer]
	cpu_status.zero_flag = register_x == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

func TXS() {
	memory.system_stack[stack_pointer] = register_x
	cpu_status.zero_flag = register_x == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

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

func CPY(val int) {
	cpu_status.carry_flag = register_y >= val
	cpu_status.zero_flag = register_y == val
	cpu_status.negative_flag = (register_y-val)&0b1000000 != 0
}

func PHP() {
	store_status := 0
	if cpu_status.carry_flag {
		store_status += 1
	}
	if cpu_status.zero_flag {
		store_status += 2
	}
	if cpu_status.interrupt_disable {
		store_status += 4
	}
	if cpu_status.decimal_mode {
		store_status += 8
	}
	if cpu_status.break_command {
		store_status += 16
	}
	if cpu_status.overflow_flag {
		store_status += 64
	}
	if cpu_status.negative_flag {
		store_status += 128
	}
	stack_pointer--
	memory.system_stack[stack_pointer] = store_status
}

func PLA() {
	stack_pointer++
	accumulator = memory.system_stack[stack_pointer]
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func PHA() {
	stack_pointer--
	memory.system_stack[stack_pointer] = accumulator
	//stack grows down!
}

func PLP() {
	stack_pointer++
	cpu_status.carry_flag = memory.system_stack[stack_pointer]&0b00000001 != 0
	cpu_status.zero_flag = memory.system_stack[stack_pointer]&0b00000010 != 0
	cpu_status.interrupt_disable = memory.system_stack[stack_pointer]&0b00000100 != 0
	cpu_status.decimal_mode = memory.system_stack[stack_pointer]&0b00001000 != 0
	cpu_status.break_command = memory.system_stack[stack_pointer]&0b00010000 != 0
	cpu_status.overflow_flag = memory.system_stack[stack_pointer]&0b01000000 != 0
	cpu_status.negative_flag = memory.system_stack[stack_pointer]&0b10000000 != 0
}

func AND(val int) {
	accumulator &= val
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

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

func BIT(val int) {
	cpu_status.zero_flag = accumulator&val == 0
	cpu_status.overflow_flag = val&0b01000000 != 0
	cpu_status.negative_flag = val&0b10000000 != 0
}

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

func DEC(mem_addr *int) {
	*mem_addr--
	cpu_status.zero_flag = *mem_addr == 0
	cpu_status.negative_flag = *mem_addr&0b10000000 != 0
}

func INC(mem_addr *int) {
	*mem_addr++
	cpu_status.zero_flag = *mem_addr == 0
	cpu_status.negative_flag = *mem_addr&0b10000000 != 0
}

func DEX() {
	register_x--
	if register_x == -1 {
		register_x = 0xff
	}
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

func LSR(addr *int) {
	cpu_status.carry_flag = accumulator&0b00000001 != 0
	if addr == nil {
		accumulator = accumulator >> 1
	} else {
		*addr = *addr >> 1
	}
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

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
		accumulator_add(-(val - 1))
	} else {
		accumulator_add(-val)
	}
	cpu_status.carry_flag = accumulator&0b100000000 != 0
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}
