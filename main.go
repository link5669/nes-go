package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

var program_counter int = 0
var cycles = 0
var stack_pointer int = 0
var accumulator int = 0
var register_x int = 0
var register_y int = 0
var cpu_status Status

var memory Memory

var const_head Const = Const{}

var head Program_Code = Program_Code{}

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
	if split_line[1] == "a" {
		ret_addr_mode = accumulator_type
		ret_value = 0
	} else {
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
	}
	return ret_addr_mode, ret_value
}

func find_const(val string) string {
	var local_ptr = &const_head
	for local_ptr.next != nil {
		val = strings.Replace(val, local_ptr.name, local_ptr.val, 1)
		local_ptr = local_ptr.next
	}
	return val
}

func read_asm() {
	f, err := os.Open(os.Args[3])
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
					split_line[1] = find_const(split_line[1])
				}
				if split_line[0] == "LSR" {
					println("alkjsd")
				}
				if strings.Contains(split_line[1], "$") || strings.Contains(split_line[1], "#") || split_line[1] == "a" {
					var input_addr_1 int
					ptr.addr_mode, input_addr_1 = get_addr_mode_val(split_line)
					switch split_line[0] {
					case "LSR":
						ptr.op_code = LSR_cmd
						ptr.mem_1 = input_addr_1
					case "ORA":
						ptr.op_code = ORA_cmd
						ptr.mem_1 = input_addr_1
					case "EOR":
						ptr.op_code = EOR_cmd
						ptr.mem_1 = input_addr_1
					case "ASL":
						ptr.op_code = ASL_cmd
						ptr.mem_1 = input_addr_1
					case "ROR":
						ptr.op_code = ROR_cmd
						ptr.mem_1 = input_addr_1
					case "ROL":
						ptr.op_code = ROL_cmd
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
					case "BVC":
						ptr.op_code = BVC_cmd
						ptr.destination = split_line[1]
					case "BMI":
						ptr.op_code = BMI_cmd
						ptr.destination = split_line[1]
					case "BVS":
						ptr.op_code = BVS_cmd
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
				case "ROR":
					ptr.op_code = ROR_cmd
				case "ROL":
					ptr.op_code = ROL_cmd
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
	dump_contents()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func reset() {
	stack_pointer = 0xFD
	cycles = 7
}

func twos_compliment(val int) int {
	//try two's compliment
	if val&0b10000000 != 0 {
		inversion := val & 0b01111111
		in_val := 0b1111111 ^ inversion
		return -1 * in_val
	} else {
		return val
	}
}

const iNESFileMagic = 0x1a53454e

type iNESFileHeader struct {
	Magic    uint32  // iNES magic number
	NumPRG   byte    // number of PRG-ROM banks (16KB each)
	NumCHR   byte    // number of CHR-ROM banks (8KB each)
	Control1 byte    // control bits
	Control2 byte    // control bits
	NumRAM   byte    // PRG-RAM size (x 8KB)
	_        [7]byte // unused padding
}

func LoadNESFile(path string) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	// read file header
	header := iNESFileHeader{}
	if err := binary.Read(file, binary.LittleEndian, &header); err != nil {
		return
	}

	// verify header magic number
	if header.Magic != iNESFileMagic {
		return
	}

	// mapper type
	// mapper1 := header.Control1 >> 4
	// mapper2 := header.Control2 >> 4
	// mapper := mapper1 | mapper2<<4

	// // mirroring type
	// mirror1 := header.Control1 & 1
	// mirror2 := (header.Control1 >> 3) & 1
	// mirror := mirror1 | mirror2<<1

	// // battery-backed RAM
	// battery := (header.Control1 >> 1) & 1

	// read trainer if present (unused)
	if header.Control1&4 == 4 {
		trainer := make([]byte, 512)
		if _, err := io.ReadFull(file, trainer); err != nil {
			return
		}
	}

	// read prg-rom bank(s)
	prg := make([]byte, int(header.NumPRG)*16384)
	if _, err := io.ReadFull(file, prg); err != nil {
		return
	}

	// read chr-rom bank(s)
	chr := make([]byte, int(header.NumCHR)*8192)
	if _, err := io.ReadFull(file, chr); err != nil {
		return
	}

	// provide chr-rom/ram if not in file
	if header.NumCHR == 0 {
		chr = make([]byte, 8192)
	}

	i := 0

	for cnt := 0xC000; cnt < 0xFABF; cnt++ {
		*memory.absolute_addr(cnt) = int(prg[cnt-0xC000])
	}

	reset()
	cpu_status.interrupt_disable = true
	program_counter = 0xC000
	for {
		if program_counter == 0xE3BD {
			dump_contents()
		}
		index := fmt.Sprintf("%x", program_counter)
		opcode := *memory.absolute_addr(program_counter)
		if run_tests(fmt.Sprint(strings.ToUpper(index), " ", "opcode", " ", "mem", " A:", strings.ToUpper(fmt.Sprintf("%x", accumulator)), " X:", strings.ToUpper(fmt.Sprintf("%x", register_x)), " Y:", strings.ToUpper(fmt.Sprintf("%x", register_y)), " P:", strings.ToUpper(fmt.Sprintf("%x", get_flag_sum())), " SP:", strings.ToUpper(fmt.Sprintf("%x", stack_pointer)), " CYC:", cycles), i) {
			break
		}
		i++
		// ptr.index = program_counter

		switch opcode {
		case 0x69:
			ADC(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0x65:
			ADC(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x75:
			ADC(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0x6D:
			ADC(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x7D:
			ADC(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0x79:
			ADC(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x61:
			ADC(*memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x71:
			ADC(*memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x29:
			AND(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0x25:
			AND(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x35:
			AND(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0x2D:
			AND(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x3D:
			AND(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0x39:
			AND(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x21:
			AND(*memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x31:
			AND(*memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x0A:
			ASL(&accumulator)
			program_counter++
			cycles += 2
		case 0x06:
			ASL(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x16:
			ASL(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x0E:
			ASL(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 6
		case 0x1E:
			ASL(memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), false))
			program_counter++
			cycles += 7
		case 0x90:
			program_counter++
			if !cpu_status.carry_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0xB0:
			program_counter++
			if cpu_status.carry_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0xF0:
			program_counter++
			if cpu_status.zero_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0x24:
			BIT(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x2C:
			BIT(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x30:
			program_counter++
			if cpu_status.negative_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0xD0:
			program_counter++
			if !cpu_status.zero_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0x10:
			program_counter++
			if !cpu_status.negative_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0x00:
			cpu_status.break_command = true
			// cycles += 7
		case 0x50:
			program_counter++
			if !cpu_status.overflow_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0x70:
			program_counter++
			if cpu_status.overflow_flag {
				cycles++
				break_to("", twos_compliment(int(prg[program_counter-0xC000]))+program_counter+1, program_counter-0xA000)
			} else {
				program_counter++
			}
			cycles += 2
		case 0x18:
			CLC()
			program_counter++
			cycles += 2
		case 0xD8:
			CLD()
			program_counter++
			cycles += 2
		case 0x58:
			CLI()
			program_counter++
			cycles += 2
		case 0xB8:
			CLV()
			program_counter++
			cycles += 2
		case 0xC9:
			CMP(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0xC5:
			CMP(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0xD5:
			CMP(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0xCD:
			CMP(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xDD:
			CMP(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0xD9:
			CMP(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xC1:
			CMP(*memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0xD1:
			CMP(*memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0xE0:
			CPX(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0xE4:
			CPX(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0xEC:
			CPX(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xC0:
			CPY(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0xC4:
			CPY(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0xCC:
			CPY(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xC6:
			DEC(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0xD6:
			DEC(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0xCE:
			DEC(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 6
		case 0xDE:
			DEC(memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), false))
			program_counter++
			cycles += 7
		case 0xCA:
			DEX()
			program_counter++
			cycles += 2
		case 0x88:
			DEY()
			program_counter++
			cycles += 2
		case 0x49:
			EOR(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0x45:
			EOR(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x55:
			EOR(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0x4D:
			EOR(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x5D:
			EOR(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0x59:
			EOR(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x41:
			EOR(*memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x51:
			EOR(*memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0xE6:
			INC(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0xF6:
			INC(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0xEE:
			INC(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 6
		case 0xFE:
			INC(memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), false))
			program_counter++
			cycles += 7
		case 0xE8:
			INX()
			program_counter++
			cycles += 2
		case 0xC8:
			INY()
			program_counter++
			cycles += 2
		case 0x4C:
			dest := double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))
			break_to("", dest, dest)
			cycles += 3
		case 0x6C:
			mem_addr := double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))
			upper_mem := mem_addr + 1
			if mem_addr&0b11111111 == 0xFF {
				upper_mem = mem_addr & 0b1111111100000000
			}
			lower := (*memory.absolute_addr(mem_addr))
			upper := (*memory.absolute_addr(upper_mem))
			loc := upper << 8
			loc += lower
			break_to("", loc, loc)
			cycles += 5
		case 0x20:
			next := program_counter + 2
			lower := next & 0b11111111
			upper := next >> 8
			memory.system_stack[stack_pointer] = upper
			stack_pointer--
			memory.system_stack[stack_pointer] = lower
			stack_pointer--
			dest := double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))
			break_to("", dest, dest)
			cycles += 6
		case 0xA9:
			LDA(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0xA5:
			LDA(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0xB5:
			LDA(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0xAD:
			LDA(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xBD:
			LDA(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0xB9:
			LDA(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xA1:
			LDA(*memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0xB1:
			LDA(*memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0xA2:
			LDX(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0xA6:
			LDX(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0xB6:
			LDX(*memory.zero_page_y_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0xAE:
			LDX(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xBE:
			LDX(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xA0:
			LDY(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0xA4:
			LDY(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0xB4:
			LDY(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0xAC:
			LDY(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xBC:
			LDY(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0x4A:
			LSR(&accumulator)
			program_counter++
			cycles += 2
		case 0x46:
			LSR(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x56:
			LSR(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x4E:
			LSR(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 6
		case 0x5E:
			LSR(memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), false))
			program_counter++
			cycles += 7
		case 0xEA:
			time.Sleep(5 * time.Microsecond)
			program_counter++
			cycles += 2
		case 0x09:
			ORA(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0x05:
			ORA(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x15:
			ORA(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0x0D:
			ORA(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x1D:
			ORA(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0x19:
			ORA(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x01:
			ORA(*memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x11:
			ORA(*memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x48:
			PHA()
			program_counter++
			cycles += 3
		case 0x08:
			PHP()
			program_counter++
			cycles += 3
		case 0x68:
			PLA()
			program_counter++
			cycles += 4
		case 0x28:
			PLP()
			program_counter++
			cycles += 4
		case 0x2A:
			ROL(&accumulator)
			program_counter++
			cycles += 2
		case 0x26:
			ROL(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x36:
			ROL(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x2E:
			ROL(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 6
		case 0x3E:
			ROL(memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), false))
			program_counter++
			cycles += 7
		case 0x6A:
			ROR(&accumulator)
			program_counter++
			cycles += 2
		case 0x66:
			ROR(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x76:
			ROR(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x6E:
			ROR(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 6
		case 0x7E:
			ROR(memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), false))
			program_counter++
			cycles += 7
		case 0x40:
			PLP()
			stack_pointer++
			lower := memory.system_stack[stack_pointer]
			stack_pointer++
			upper := memory.system_stack[stack_pointer]
			last_index := upper << 8
			last_index += lower
			program_counter = last_index
			cycles += 6
		case 0x60:
			stack_pointer++
			lower := memory.system_stack[stack_pointer]
			stack_pointer++
			upper := memory.system_stack[stack_pointer]
			last_index := upper << 8
			last_index += lower
			// last_index -= 2
			program_counter = last_index
			program_counter++
			cycles += 6
		case 0xE9:
			SBC(single_addr(byte(*memory.absolute_addr(program_counter + 1))))
			program_counter++
			cycles += 2
		case 0xE5:
			SBC(*memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0xF5:
			SBC(*memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0xED:
			SBC(*memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xFD:
			SBC(*memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), true))
			program_counter++
			cycles += 4
		case 0xF9:
			SBC(*memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xE1:
			SBC(*memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0xF1:
			SBC(*memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 5
		case 0x38:
			SEC()
			program_counter++
			cycles += 2
		case 0xF8:
			SED()
			program_counter++
			cycles += 2
		case 0x78:
			SEI()
			program_counter++
			cycles += 2
		case 0x85:
			STA(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x95:
			STA(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0x8D:
			STA(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x9D:
			STA(memory.absolute_x_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2))), false))
			program_counter++
			cycles += 5
		case 0x99:
			STA(memory.absolute_y_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 5
		case 0x81:
			STA(memory.indexed_indirect_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x91:
			STA(memory.indirect_indexed_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 6
		case 0x86:
			STX(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x96:
			STX(memory.zero_page_y_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0x8E:
			STX(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0x84:
			STY(memory.zero_page_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 3
		case 0x94:
			STY(memory.zero_page_x_addr(single_addr(byte(*memory.absolute_addr(program_counter + 1)))))
			program_counter++
			cycles += 4
		case 0x8C:
			STY(memory.absolute_addr(double_addr(byte(*memory.absolute_addr(program_counter + 1)), byte(*memory.absolute_addr(program_counter + 2)))))
			program_counter++
			cycles += 4
		case 0xAA:
			TAX()
			program_counter++
			cycles += 2
		case 0xA8:
			TAY()
			program_counter++
			cycles += 2
		case 0xBA:
			TSX()
			program_counter++
			cycles += 2
		case 0x8A:
			TXA()
			program_counter++
			cycles += 2
		case 0x9A:
			TXS()
			program_counter++
			cycles += 2
		case 0x98:
			TYA()
			program_counter++
			cycles += 2
			//ILLEGAL OPCODES
			// case 0x67:
			// 	ptr.op_code = RRA_cmd
			// 	cycles += 5
			// 	ptr.addr_mode = zero_page_type
			// 	ptr.mem_1 = single_addr(byte(*memory.absolute_addr(program_counter+1)))
			// case 0x77:
			// 	ptr.op_code = RRA_cmd
			// 	cycles += 6
			// 	ptr.addr_mode = zero_page_x_type
			// 	ptr.mem_1 = single_addr(byte(*memory.absolute_addr(program_counter+1)))
			// case 0x6F:
			// 	ptr.op_code = RRA_cmd
			// 	cycles += 6
			// 	ptr.addr_mode = absolute_type
			// 	ptr.mem_1 = double_addr(byte(*memory.absolute_addr(program_counter+1)), byte(*memory.absolute_addr(program_counter+2)))
			// case 0x7F:
			// 	ptr.op_code = RRA_cmd
			// 	ptr.cycles = 7
			// 	ptr.addr_mode = absolute_x_type
			// 	ptr.mem_1 = double_addr(byte(*memory.absolute_addr(program_counter+1)), byte(*memory.absolute_addr(program_counter+2)))
			// case 0x7B:
			// 	ptr.op_code = RRA_cmd
			// 	ptr.cycles = 7
			// 	ptr.addr_mode = absolute_y_type
			// 	ptr.mem_1 = double_addr(byte(*memory.absolute_addr(program_counter+1)), byte(*memory.absolute_addr(program_counter+2)))
			// case 0x63:
			// 	ptr.op_code = RRA_cmd
			// 	ptr.cycles = 8
			// 	ptr.addr_mode = indexed_indirect_type
			// 	ptr.mem_1 = single_addr(byte(*memory.absolute_addr(program_counter+1)))
			// case 0x73:
			// 	ptr.op_code = RRA_cmd
			// 	ptr.cycles = 8
			// 	ptr.addr_mode = indirect_indexed_type
			// 	ptr.mem_1 = single_addr(byte(*memory.absolute_addr(program_counter+1)))
			// case 0x87:
			// 	ptr.op_code = SAX_cmd
			// 	ptr.addr_mode = zero_page_type
			// 	ptr.mem_1 = single_addr(byte(*memory.absolute_addr(program_counter+1)))
			// 	cycles += 3
			// case 0x97:
			// 	ptr.op_code = SAX_cmd
			// 	ptr.addr_mode = zero_page_y_type
			// 	ptr.mem_1 = single_addr(byte(*memory.absolute_addr(program_counter+1)))
			// 	cycles += 4
			// case 0x8F:
			// 	ptr.op_code = SAX_cmd
			// 	ptr.addr_mode = absolute_type
			// 	ptr.mem_1 = double_addr(byte(*memory.absolute_addr(program_counter+1)), byte(*memory.absolute_addr(program_counter+2)))
			// 	cycles += 4
			// case 0x83:
			// 	ptr.op_code = SAX_cmd
			// 	ptr.addr_mode = indexed_indirect_type
			// 	ptr.mem_1 = single_addr(byte(*memory.absolute_addr(program_counter+1)))
			// 	cycles += 6
		}
	}
	// dump_program()
	dump_contents()
}

func single_addr(val byte) int {
	program_counter += 1
	// memory.gen_memory[program_counter-0xA000] = int(val)
	return int(val)
}

func double_addr(first byte, second byte) int {
	program_counter++
	memory.gen_memory[program_counter-0xA000] = int(first)
	program_counter++
	memory.gen_memory[program_counter-0xA000] = int(second)
	lower := fmt.Sprintf("%x", int(first))
	upper := fmt.Sprintf("%x", int(second))
	if first < 0x10 {
		lower = "0" + lower
	}
	if second < 0x10 {
		upper = "0" + upper
	}
	ret_val := string_to_int("$" + upper + lower)
	return ret_val
}

func main() {
	if os.Args[1] == "--input" {
		if os.Args[2] == "asm" {
			read_asm()
		} else if os.Args[2] == "rom" {
			LoadNESFile(os.Args[3])
		}
	}
}

func get_flag_sum() int {
	flags := 0
	if cpu_status.carry_flag {
		flags += 1
	}
	if cpu_status.zero_flag {
		flags += 2
	}
	if cpu_status.interrupt_disable {
		flags += 4
	}
	if cpu_status.decimal_mode {
		flags += 8
	}
	if cpu_status.break_command {
		flags += 16
	}
	flags += 32
	if cpu_status.overflow_flag {
		flags += 64
	}
	if cpu_status.negative_flag {
		flags += 128
	}
	return flags
}

func break_to(destination string, mem_dest int, mem_start int) {
	program_counter = mem_dest
	//FIX THIS
	var m256a int = (mem_start - head.index) + 256 - ((mem_start - head.index) % 256)
	var m256b int = mem_dest + 256 - (mem_dest % 256)
	if (m256a <= (mem_start-head.index) && (mem_start-head.index) < m256b) || (m256b <= mem_dest && mem_dest < m256a) {
		cycles += 2
	}
}

func SAX(mem_addr *int) {
	*mem_addr = accumulator & register_x
}

func load(val int, desti *int) {
	*desti = val
	cpu_status.zero_flag = *desti == 0
	cpu_status.negative_flag = *desti&0b10000000 != 0
}

func LDA(val int) {
	load(val, &accumulator)
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
	//need to wrap?
	*mem_addr = register_x
}

func STY(mem_addr *int) {
	//need to wrap?
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
	cpu_status.negative_flag = register_y&0b10000000 != 0
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
	register_x = stack_pointer
	cpu_status.zero_flag = register_x == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

func TXS() {
	stack_pointer = register_x
}

func CPX(val int) {
	cpu_status.carry_flag = register_x >= val
	cpu_status.zero_flag = register_x == val
	cpu_status.negative_flag = (register_x-val)&0b10000000 != 0
}

func CMP(val int) {
	cpu_status.carry_flag = accumulator >= val
	cpu_status.zero_flag = accumulator == val
	cpu_status.negative_flag = (accumulator-val)&0b10000000 != 0
}

func CPY(val int) {
	cpu_status.carry_flag = register_y >= val
	cpu_status.zero_flag = register_y == val
	cpu_status.negative_flag = (register_y-val)&0b10000000 != 0
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
	//BRK is always true when pushed from software
	store_status += 16
	store_status += 32
	if cpu_status.overflow_flag {
		store_status += 64
	}
	if cpu_status.negative_flag {
		store_status += 128
	}
	memory.system_stack[stack_pointer] = store_status
	stack_pointer--
}

func PLA() {
	stack_pointer++
	accumulator = memory.system_stack[stack_pointer]
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func PHA() {
	memory.system_stack[stack_pointer] = accumulator
	stack_pointer--
}

func PLP() {
	stack_pointer++
	cpu_status.carry_flag = memory.system_stack[stack_pointer]&0b00000001 != 0
	cpu_status.zero_flag = memory.system_stack[stack_pointer]&0b00000010 != 0
	cpu_status.interrupt_disable = memory.system_stack[stack_pointer]&0b00000100 != 0
	cpu_status.decimal_mode = memory.system_stack[stack_pointer]&0b00001000 != 0
	cpu_status.break_command = false
	cpu_status.overflow_flag = memory.system_stack[stack_pointer]&0b01000000 != 0
	cpu_status.negative_flag = memory.system_stack[stack_pointer]&0b10000000 != 0
}

func AND(val int) {
	accumulator &= val
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func EOR(mem_addr int) {
	accumulator ^= mem_addr
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func ORA(mem_addr int) {
	accumulator |= mem_addr
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func BIT(val int) {
	cpu_status.zero_flag = accumulator&val == 0
	cpu_status.overflow_flag = val&0b01000000 != 0
	cpu_status.negative_flag = val&0b10000000 != 0
}

func INX() {
	register_x++
	if register_x > 255 {
		register_x -= 256
	}
	cpu_status.zero_flag = register_x == 0
	cpu_status.negative_flag = register_x&0b10000000 != 0
}

func INY() {
	register_y++
	if register_y > 255 {
		register_y -= 256
	}
	cpu_status.zero_flag = register_y == 0
	cpu_status.negative_flag = register_y&0b10000000 != 0
}

func DEC(mem_addr *int) {
	*mem_addr--
	if *mem_addr < 0 {
		*mem_addr += 256
	}
	cpu_status.zero_flag = *mem_addr == 0
	cpu_status.negative_flag = *mem_addr&0b10000000 != 0
}

func INC(mem_addr *int) {
	*mem_addr++
	if *mem_addr > 255 {
		*mem_addr -= 256
	}
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
	if register_y == -1 {
		register_y = 0xff
	}
	cpu_status.zero_flag = register_y == 0
	cpu_status.negative_flag = register_y&0b10000000 != 0
}

func ASL(mem_addr *int) {
	cpu_status.carry_flag = *mem_addr&0b10000000 != 0
	*mem_addr = *mem_addr << 1
	if *mem_addr > 255 {
		*mem_addr -= 256
	}
	cpu_status.zero_flag = *mem_addr == 0
	cpu_status.negative_flag = *mem_addr&0b10000000 != 0
}

func LSR(addr *int) {
	cpu_status.carry_flag = *addr&0b00000001 != 0
	*addr = *addr >> 1
	cpu_status.zero_flag = *addr == 0
	cpu_status.negative_flag = *addr&0b10000000 != 0
}

func ROR(addr *int) {
	temp := *addr
	*addr = *addr >> 1
	if cpu_status.carry_flag {
		*addr |= 0b10000000
	} else {
		*addr &= 0b01111111
	}
	if *addr > 255 {
		*addr -= 256
	}
	cpu_status.carry_flag = 0b00000001&temp != 0
	cpu_status.zero_flag = *addr == 0
	cpu_status.negative_flag = 0b10000000&*addr != 0
}

func ROL(addr *int) {
	temp := *addr
	*addr = *addr << 1
	if cpu_status.carry_flag {
		*addr |= 0b00000001
	} else {
		*addr &= 0b11111110
	}
	if *addr > 255 {
		*addr -= 256
	}
	cpu_status.carry_flag = 0b10000000&temp != 0
	cpu_status.zero_flag = *addr == 0
	cpu_status.negative_flag = 0b10000000&*addr != 0
}

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
		cpu_status.carry_flag = accumulator+val+1 > 0xFF
		cpu_status.overflow_flag = (accumulator^val)&0x80 == 0 && (accumulator^accumulator+val+1)&0x80 != 0
		accumulator += val + 1
	} else {
		cpu_status.carry_flag = accumulator+val > 0xFF
		cpu_status.overflow_flag = (accumulator^val)&0x80 == 0 && (accumulator^accumulator+val)&0x80 != 0
		accumulator += val
	}
	if cpu_status.carry_flag {
		accumulator -= 0x100
	}
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func SBC(val int) {
	if !cpu_status.carry_flag {
		cpu_status.carry_flag = !(accumulator-val < 0)
		cpu_status.overflow_flag = (accumulator^val)&0x80 != 0 && (accumulator^accumulator-val-1)&0x80 != 0
		accumulator -= val + 1
	} else {
		cpu_status.carry_flag = !(accumulator-val < 0)
		cpu_status.overflow_flag = (accumulator^val)&0x80 != 0 && (accumulator^accumulator-val)&0x80 != 0
		accumulator -= val
	}
	if accumulator < 0 {
		accumulator += 256
	}
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}
