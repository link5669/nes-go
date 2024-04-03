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

func string_to_hex(str string) int {
	split := strings.TrimPrefix(str, "#")
	split = strings.TrimPrefix(split, "$")
	split = strings.TrimLeft(split, "0")
	if len(split) > 1 {
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
	for scanner.Scan() {
		curr_line := scanner.Text()
		split_line := strings.Split(curr_line, " ")
		switch split_line[0] {
		case "LDA":
			if strings.Contains(split_line[1], "#") {
				LDA(string_to_hex(split_line[1]))
			}
		case "STA":
			STA(memory.zero_page_addr(string_to_hex(split_line[1])))
		}
	}

	dump_contents()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

// func STX(mem_addr int) {
// 	memory[mem_addr] = register_x
// }

// func STY(mem_addr int) {
// 	memory[mem_addr] = register_y
// }

// func TAX() {
// 	register_x = accumulator
// 	cpu_status.zero_flag = register_x == 0
// 	cpu_status.negative_flag = register_x&0b10000000 != 0

// }

// func TAY() {
// 	register_y = accumulator
// 	cpu_status.zero_flag = register_y == 0
// 	cpu_status.negative_flag = register_x&0b10000000 != 0

// }

// func TXA() {
// 	accumulator = register_x
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0
// }

// func TYA() {
// 	accumulator = register_y
// 	cpu_status.zero_flag = accumulator == 0
// 	cpu_status.negative_flag = accumulator&0b10000000 != 0

// }

// func TSX() {
// 	register_x = memory[stack_pointer]
// 	cpu_status.zero_flag = register_x == 0
// 	cpu_status.negative_flag = register_x&0b10000000 != 0
// }

// func TXS() {
// 	memory[stack_pointer] = register_x
// }

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

// func INX() {
// 	register_x++
// 	cpu_status.zero_flag = register_x == 0
// 	cpu_status.negative_flag = register_x&0b10000000 != 0
// }

// func INY() {
// 	register_y++
// 	cpu_status.zero_flag = register_y == 0
// 	cpu_status.negative_flag = register_y&0b10000000 != 0
// }

// func DEC(mem_addr int) {
// 	memory[mem_addr]--
// 	cpu_status.zero_flag = memory[mem_addr] == 0
// 	cpu_status.negative_flag = memory[mem_addr]&0b10000000 != 0
// }

// func DEX() {
// 	register_x++
// 	cpu_status.zero_flag = register_x == 0
// 	cpu_status.negative_flag = register_x&0b10000000 != 0
// }

// func DEY() {
// 	register_y++
// 	cpu_status.zero_flag = register_y == 0
// 	cpu_status.negative_flag = register_y&0b10000000 != 0
// }

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

// func CLC() {
// 	cpu_status.carry_flag = false
// }

// func CLD() {
// 	cpu_status.decimal_mode = false
// }

// func CLI() {
// 	cpu_status.interrupt_disable = false
// }

// func CLV() {
// 	cpu_status.overflow_flag = false
// }

// func SEC() {
// 	cpu_status.carry_flag = true
// }

// func SED() {
// 	cpu_status.decimal_mode = true
// }

// func SEI() {
// 	cpu_status.interrupt_disable = true
// }

func dump_contents() {
	fmt.Println("Accumulator:", accumulator)
	fmt.Println("X Register:", register_x)
	fmt.Println("Y Register:", register_y)
	fmt.Println("Stack Pointer:", stack_pointer)
	fmt.Println("Program Counter:", program_counter)

	fmt.Println("Zero Page:\n")
	for i := 0; i < 16; i++ {
		fmt.Print(i, ": ")
		for j := 0; j < 16; j++ {
			fmt.Print(memory.zero_page[i*16+j])
		}
		fmt.Print("\n")
	}
}
