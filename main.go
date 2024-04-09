// https://skilldrick.github.io/easy6502/ -- use for testing!
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
				if strings.Contains(split_line[1], "$") || strings.Contains(split_line[1], "#") {
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
	run_program()
	dump_contents()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func reset() {
	stack_pointer = 0xFD
	cycles = 7
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

	var ptr *Program_Code
	ptr = &head
	// var var_ptr *Const
	// var_ptr = &const_head
	for program_counter = 0xC000; program_counter < 0xFF97; program_counter++ {
		ptr.index = program_counter
		opcode := prg[program_counter-0xC000]
		//assume all are opcodes
		ptr.code_type = op_code
		memory.gen_memory[program_counter-0x6000] = int(opcode)
		ptr.addr_in_mem = program_counter - 0x6000
		switch opcode {
		case 0x69:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0x65:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x75:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0x6D:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x7D:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x79:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x61:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x71:
			ptr.op_code = ADC_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x29:
			ptr.op_code = AND_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0x25:
			ptr.op_code = AND_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x35:
			ptr.op_code = AND_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0x2D:
			ptr.op_code = AND_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x3D:
			ptr.op_code = AND_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x39:
			ptr.op_code = AND_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x21:
			ptr.op_code = AND_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x31:
			ptr.op_code = AND_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x0A:
			ptr.op_code = ASL_cmd
			ptr.addr_mode = accumulator_type
			ptr.cycles = 2
		case 0x06:
			ptr.op_code = ASL_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x16:
			ptr.op_code = ASL_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x0E:
			ptr.op_code = ASL_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 6
		case 0x1E:
			ptr.op_code = ASL_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 7
		case 0x90:
			ptr.op_code = BCC_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0xB0:
			ptr.op_code = BCS_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0xF0:
			ptr.op_code = BEQ_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0x24:
			ptr.op_code = BIT_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x2C:
			ptr.op_code = BIT_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x30:
			ptr.op_code = BMI_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0xD0:
			ptr.op_code = BNE_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0x10:
			ptr.op_code = BPL_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0x00:
			ptr.op_code = BRK_cmd
			ptr.cycles = 7
		case 0x50:
			ptr.op_code = BVC_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0x70:
			ptr.op_code = BVS_cmd
			program_counter++
			ptr.mem_1 = int(prg[program_counter-0xC000]) + ptr.index + 2
			ptr.cycles = 2
		case 0x18:
			ptr.op_code = CLC_cmd
			ptr.cycles = 2
		case 0xD8:
			ptr.op_code = CLD_cmd
			ptr.cycles = 2
		case 0x58:
			ptr.op_code = CLI_cmd
			ptr.cycles = 2
		case 0xB8:
			ptr.op_code = CLV_cmd
			ptr.cycles = 2
		case 0xC9:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0xC5:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0xD5:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0xCD:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xDD:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xD9:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xC1:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0xD1:
			ptr.op_code = CMP_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0xE0:
			ptr.op_code = CPX_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0xE4:
			ptr.op_code = CPX_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0xEC:
			ptr.op_code = CPX_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xC0:
			ptr.op_code = CPY_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0xC4:
			ptr.op_code = CPY_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0xCC:
			ptr.op_code = CPY_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xC6:
			ptr.op_code = DEC_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0xD6:
			ptr.op_code = DEC_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0xCE:
			ptr.op_code = DEC_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 6
		case 0xDE:
			ptr.op_code = DEC_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 7
		case 0xCA:
			ptr.op_code = DEX_cmd
			ptr.cycles = 2
		case 0x88:
			ptr.op_code = DEY_cmd
			ptr.cycles = 2
		case 0x49:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0x45:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x55:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0x4D:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x5D:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x59:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x41:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x51:
			ptr.op_code = EOR_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0xE6:
			ptr.op_code = INC_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0xF6:
			ptr.op_code = INC_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0xEE:
			ptr.op_code = INC_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 6
		case 0xFE:
			ptr.op_code = INC_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 7
		case 0xE8:
			ptr.op_code = INX_cmd
			ptr.cycles = 2
		case 0xC8:
			ptr.op_code = INY_cmd
			ptr.cycles = 2
		case 0x4C:
			ptr.op_code = JMP_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 3
		case 0x6C:
			ptr.op_code = JMP_cmd
			ptr.addr_mode = indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x20:
			ptr.op_code = JSR_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 6
		case 0xA9:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0xA5:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0xB5:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0xAD:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xBD:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xB9:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xA1:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0xB1:
			ptr.op_code = LDA_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0xA2:
			ptr.op_code = LDX_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0xA6:
			ptr.op_code = LDX_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0xB6:
			ptr.op_code = LDX_cmd
			ptr.addr_mode = zero_page_y_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0xAE:
			ptr.op_code = LDX_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xBE:
			ptr.op_code = LDX_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xA0:
			ptr.op_code = LDY_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0xA4:
			ptr.op_code = LDY_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0xB4:
			ptr.op_code = LDY_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0xAC:
			ptr.op_code = LDY_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xBC:
			ptr.op_code = LDY_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x4A:
			ptr.op_code = LSR_cmd
			ptr.addr_mode = implicit_type
			ptr.cycles = 2
		case 0x46:
			ptr.op_code = LSR_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x56:
			ptr.op_code = LSR_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x4E:
			ptr.op_code = LSR_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 6
		case 0x5E:
			ptr.op_code = LSR_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 7
		case 0xEA:
			ptr.op_code = NOP_cmd
			ptr.cycles = 2
		case 0x09:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0x05:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x15:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0x0D:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x1D:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x19:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x01:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x11:
			ptr.op_code = ORA_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x48:
			ptr.op_code = PHA_cmd
			ptr.cycles = 3
		case 0x08:
			ptr.op_code = PHP_cmd
			ptr.cycles = 3
		case 0x68:
			ptr.op_code = PLA_cmd
			ptr.cycles = 4
		case 0x28:
			ptr.op_code = PLP_cmd
			ptr.cycles = 4
		case 0x2A:
			ptr.op_code = ROL_cmd
			ptr.addr_mode = implicit_type
			ptr.cycles = 2
		case 0x26:
			ptr.op_code = ROL_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x36:
			ptr.op_code = ROL_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x2E:
			ptr.op_code = ROL_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 6
		case 0x3E:
			ptr.op_code = ROL_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 7
		case 0x6A:
			ptr.op_code = ROR_cmd
			ptr.addr_mode = implicit_type
			ptr.cycles = 2
		case 0x66:
			ptr.op_code = ROR_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x76:
			ptr.op_code = ROR_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x6E:
			ptr.op_code = ROR_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 6
		case 0x7E:
			ptr.op_code = ROR_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 7
		case 0x40:
			ptr.op_code = RTI_cmd
			println("IMPLEMENT RTI")
			ptr.cycles = 6
		case 0x60:
			ptr.op_code = RTS_cmd
			ptr.cycles = 6
		case 0xE9:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = immediate_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 2
		case 0xE5:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0xF5:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0xED:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xFD:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xF9:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xE1:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0xF1:
			ptr.op_code = SBC_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 5
		case 0x38:
			ptr.op_code = SEC_cmd
			ptr.cycles = 2
		case 0xF8:
			ptr.op_code = SED_cmd
			ptr.cycles = 2
		case 0x78:
			ptr.op_code = SEI_cmd
			ptr.cycles = 2
		case 0x85:
			ptr.op_code = STA_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x95:
			ptr.op_code = STA_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0x8D:
			ptr.op_code = STA_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x9D:
			ptr.op_code = STA_cmd
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 5
		case 0x99:
			ptr.op_code = STA_cmd
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 5
		case 0x81:
			ptr.op_code = STA_cmd
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x91:
			ptr.op_code = STA_cmd
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 6
		case 0x86:
			ptr.op_code = STX_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x96:
			ptr.op_code = STX_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0x8E:
			ptr.op_code = STX_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0x84:
			ptr.op_code = STY_cmd
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 3
		case 0x94:
			ptr.op_code = STY_cmd
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
			ptr.cycles = 4
		case 0x8C:
			ptr.op_code = STY_cmd
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
			ptr.cycles = 4
		case 0xAA:
			ptr.op_code = TAX_cmd
			ptr.cycles = 2
		case 0xA8:
			ptr.op_code = TAY_cmd
			ptr.cycles = 2
		case 0xBA:
			ptr.op_code = TSX_cmd
			ptr.cycles = 2
		case 0x8A:
			ptr.op_code = TXA_cmd
			ptr.cycles = 2
		case 0x9A:
			ptr.op_code = TXS_cmd
			ptr.cycles = 2
		case 0x98:
			ptr.op_code = TYA_cmd
			ptr.cycles = 2
		//ILLEGAL OPCODES
		case 0x67:
			ptr.op_code = RRA_cmd
			ptr.cycles = 5
			ptr.addr_mode = zero_page_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
		case 0x77:
			ptr.op_code = RRA_cmd
			ptr.cycles = 6
			ptr.addr_mode = zero_page_x_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
		case 0x6F:
			ptr.op_code = RRA_cmd
			ptr.cycles = 6
			ptr.addr_mode = absolute_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
		case 0x7F:
			ptr.op_code = RRA_cmd
			ptr.cycles = 7
			ptr.addr_mode = absolute_x_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
		case 0x7B:
			ptr.op_code = RRA_cmd
			ptr.cycles = 7
			ptr.addr_mode = absolute_y_type
			ptr.mem_1 = double_addr(prg[program_counter-0xC000+1], prg[program_counter-0xC000+2])
		case 0x63:
			ptr.op_code = RRA_cmd
			ptr.cycles = 8
			ptr.addr_mode = indexed_indirect_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
		case 0x73:
			ptr.op_code = RRA_cmd
			ptr.cycles = 8
			ptr.addr_mode = indirect_indexed_type
			ptr.mem_1 = single_addr(prg[program_counter-0xC000+1])
		}
		ptr.next = &Program_Code{}
		ptr = ptr.next
	}
	// dump_program()
	run_program()
	// dump_contents()
}

func single_addr(val byte) int {
	program_counter += 1
	memory.gen_memory[program_counter-0x6000] = int(val)
	return int(val)
}

func double_addr(first byte, second byte) int {
	program_counter++
	memory.gen_memory[program_counter-0x6000] = int(first)
	program_counter++
	memory.gen_memory[program_counter-0x6000] = int(second)
	lower := fmt.Sprintf("%x", int(first))
	upper := fmt.Sprintf("%x", int(second))
	if first < 0x10 {
		lower = "0" + lower
	}
	if second < 0x10 {
		upper = "0" + upper
	}
	return string_to_int("$" + upper + lower)
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

func run_program() {
	reset()
	var ptr *Program_Code = &head
	cpu_status.interrupt_disable = true
	i := 0
	for ptr.next != nil {
		// if ptr.index == 0xc75D {
		// 	println("alksd")
		// }
		index := fmt.Sprintf("%x", ptr.index)
		mem := fmt.Sprintf("%x", ptr.mem_1)
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
		run_tests(fmt.Sprint(strings.ToUpper(index), " ", printOpCode(ptr.op_code), " ", mem, " A:", strings.ToUpper(fmt.Sprintf("%x", accumulator)), " X:", strings.ToUpper(fmt.Sprintf("%x", register_x)), " Y:", strings.ToUpper(fmt.Sprintf("%x", register_y)), " P:", strings.ToUpper(fmt.Sprintf("%x", flags)), " SP:", strings.ToUpper(fmt.Sprintf("%x", stack_pointer)), " CYC:", cycles), i)
		i++
		if ptr.index >= 0xC949 {
			break
		}
		cycles += ptr.cycles
		if ptr.code_type == function_definition {
			//TODO
		} else {
			//for real vals
			var val int
			var val_ptr *int
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
			if ptr.addr_mode == zero_page_type {
				val_ptr = (memory.zero_page_addr(ptr.mem_1))
			} else if ptr.addr_mode == zero_page_x_type {
				val_ptr = (memory.zero_page_x_addr(ptr.mem_1))
			} else if ptr.addr_mode == zero_page_y_type {
				val_ptr = (memory.zero_page_y_addr(ptr.mem_1))
			} else if ptr.addr_mode == absolute_type {
				val_ptr = (memory.absolute_addr(ptr.mem_1))
			} else if ptr.addr_mode == absolute_x_type {
				val_ptr = (memory.absolute_x_addr(ptr.mem_1))
			} else if ptr.addr_mode == absolute_y_type {
				val_ptr = (memory.absolute_y_addr(ptr.mem_1))
			} else if ptr.addr_mode == indexed_indirect_type {
				val_ptr = (memory.indexed_indirect_addr(ptr.mem_1))
			} else if ptr.addr_mode == indirect_indexed_type {
				val_ptr = (memory.indirect_indexed_addr(ptr.mem_1))
			} else if ptr.addr_mode == accumulator_type {
				val_ptr = &accumulator
			}
			switch ptr.op_code {
			case DEC_cmd:
				DEC(val_ptr)
			case BIT_cmd:
				BIT(val)
			case INC_cmd:
				INC(val_ptr)
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
			case ORA_cmd:
				ORA(val)
			case ASL_cmd:
				ASL(val_ptr)
			case EOR_cmd:
				EOR(val)
			case CMP_cmd:
				CMP(val)
			case CPY_cmd:
				CPY(val)
			case STA_cmd:
				STA(val_ptr)
			case AND_cmd:
				AND(val)
			case LSR_cmd:
				//fix me!
				if ptr.addr_mode == implicit_type {
					LSR(nil)
				} else {
					LSR(val_ptr)
				}
			case ROR_cmd:
				if ptr.addr_mode == implicit_type {
					ROR(&accumulator)
				} else {
					ROR(val_ptr)
				}
			case ROL_cmd:
				if ptr.addr_mode == implicit_type {
					ROR(&accumulator)
				} else {
					ROR(val_ptr)
				}
			case PLP_cmd:
				PLP()
			case PHP_cmd:
				PHP()
			case BRK_cmd:
				cpu_status.break_command = true
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
				STX(val_ptr)
			case STY_cmd:
				STY(val_ptr)
			case ADC_cmd:
				ADC(val)
			case BNE_cmd:
				if !cpu_status.zero_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case BMI_cmd:
				if cpu_status.negative_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case BEQ_cmd:
				if cpu_status.zero_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case BVS_cmd:
				if cpu_status.overflow_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case BVC_cmd:
				if !cpu_status.overflow_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case BCS_cmd:
				if cpu_status.carry_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case BCC_cmd:
				if !cpu_status.carry_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case BPL_cmd:
				if !cpu_status.negative_flag {
					cycles++
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.addr_in_mem)
					continue
				}
			case JMP_cmd:
				if ptr.addr_mode == indirect_type {
					if ptr.destination != "" {
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
					} else {
						println("implement indirect jmp for hex!")
					}
				} else if ptr.addr_mode == absolute_type {
					ptr = break_to(ptr.destination, ptr.mem_1, ptr.mem_1)
					continue
				}
			case JSR_cmd:
				stack_pointer--
				lower := ptr.index & 0b11111111
				upper := ptr.index >> 8
				memory.system_stack[stack_pointer] = lower
				stack_pointer--
				memory.system_stack[stack_pointer] = upper
				ptr = break_to(ptr.destination, ptr.mem_1, ptr.mem_1)
				continue
			case RTS_cmd:
				upper := memory.system_stack[stack_pointer]
				stack_pointer++
				lower := memory.system_stack[stack_pointer]
				stack_pointer++
				last_index := upper << 8
				last_index += lower
				var local_ptr = &head
				for local_ptr.next != nil {
					if local_ptr.index == last_index {
						ptr = local_ptr
						break
					}
					local_ptr = local_ptr.next
				}
			}
		}
		ptr = ptr.next
	}
}

func break_to(destination string, mem_dest int, mem_start int) *Program_Code {
	var local_ptr = &head
	for local_ptr.next != nil {
		if (local_ptr.destination == destination && local_ptr.code_type == function_definition) || local_ptr.index == mem_dest {
			//FIX THISS
			var m256a int = mem_start + 256 - (mem_start % 256)
			var m256b int = mem_dest + 256 - (mem_dest % 256)
			if (m256a <= mem_start && mem_start < m256b) || (m256b <= mem_dest && mem_dest < m256a) {
				cycles += 2
			}
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
	//BRK is always true when pushed from software
	store_status += 16
	store_status += 32
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
	accumulator = memory.system_stack[stack_pointer]
	stack_pointer++
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func PHA() {
	stack_pointer--
	memory.system_stack[stack_pointer] = accumulator
}

func PLP() {
	cpu_status.carry_flag = memory.system_stack[stack_pointer]&0b00000001 != 0
	cpu_status.zero_flag = memory.system_stack[stack_pointer]&0b00000010 != 0
	cpu_status.interrupt_disable = memory.system_stack[stack_pointer]&0b00000100 != 0
	cpu_status.decimal_mode = memory.system_stack[stack_pointer]&0b00001000 != 0
	cpu_status.break_command = false
	cpu_status.overflow_flag = memory.system_stack[stack_pointer]&0b01000000 != 0
	cpu_status.negative_flag = memory.system_stack[stack_pointer]&0b10000000 != 0
	stack_pointer++
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

func ASL(mem_addr *int) {
	cpu_status.carry_flag = *mem_addr&0b10000000 != 0
	*mem_addr = *mem_addr << 1
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = *mem_addr&0b10000000 != 0
}

func LSR(addr *int) {
	cpu_status.carry_flag = accumulator&0b00000001 != 0
	if addr == nil {
		accumulator = accumulator >> 1
	} else {
		*addr = *addr >> 1
	}
	cpu_status.negative_flag = accumulator&0b10000000 != 0
}

func ROR(addr *int) {
	temp := *addr
	*addr = *addr >> 1
	if cpu_status.carry_flag {
		*addr |= 0b11111111
	} else {
		*addr &= 0b01111111
	}
	cpu_status.carry_flag = 0b00000001&temp != 0
	cpu_status.zero_flag = accumulator == 0
	cpu_status.negative_flag = 0b10000000&*addr != 0
}

func ROL(addr *int) {
	temp := *addr
	*addr = *addr << 1
	if cpu_status.carry_flag {
		*addr |= 0b11111111
	} else {
		*addr &= 0b01111111
	}
	cpu_status.carry_flag = 0b00000001&temp != 0
	cpu_status.zero_flag = accumulator == 0
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
