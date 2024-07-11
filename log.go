package main

import "fmt"

func dump_contents(cpu *CPU) {
	fmt.Println("Accumulator:", cpu.accumulator)
	fmt.Println("X Register:", cpu.register_x)
	fmt.Println("Y Register:", cpu.register_y)
	fmt.Println("Stack Pointer:", cpu.stack_pointer)
	fmt.Println("Program Counter:", cpu.program_counter)
	fmt.Println("Cycles:", cpu.cycles)
	fmt.Println("\nFlag Status:")
	fmt.Println("Program Counter:", cpu.cpu_status.carry_flag)
	fmt.Println("Zero Flag:", cpu.cpu_status.zero_flag)
	fmt.Println("Interrupt Disable:", cpu.cpu_status.interrupt_disable)
	fmt.Println("Decimal Mode:", cpu.cpu_status.decimal_mode)
	fmt.Println("Break Command:", cpu.cpu_status.break_command)
	fmt.Println("Overflow Flag:", cpu.cpu_status.overflow_flag)
	fmt.Println("Negative Flag:", cpu.cpu_status.negative_flag)

	fmt.Println("\nZero Page:")
	for i := 0; i < 16; i++ {
		fmt.Print(i, ": ")
		for j := 0; j < 16; j++ {
			fmt.Print(memory.zero_page[i*16+j], " ")
		}
		fmt.Print("\n")
	}

	fmt.Println("\nSystem Stack:")
	for i := 0; i < 16; i++ {
		fmt.Print(i, ": ")
		for j := 0; j < 16; j++ {
			fmt.Print(memory.system_stack[i*16+j], " ")
		}
		fmt.Print("\n")
	}
}

func print_screen() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			if memory.gen_memory[i*32+j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(memory.gen_memory[i*32+j])
			}
		}
		fmt.Print("\n")
	}
}

func printAddrModeType(ct addr_mode) string {
	switch ct {
	case implicit_type:
		return "Implicit Type"
	case accumulator_type:
		return "Accumulator Type"
	case immediate_type:
		return "Immediate Type"
	case zero_page_type:
		return "Zero Page Type"
	case zero_page_x_type:
		return "Zero Page X Type"
	case zero_page_y_type:
		return "Zero Page Y Type"
	case relative_type:
		return "Relative Type"
	case absolute_type:
		return "Absolute Type"
	case absolute_x_type:
		return "Absolute X Type"
	case absolute_y_type:
		return "Absolute Y Type"
	case indirect_type:
		return "Indirect Type"
	case indexed_indirect_type:
		return "Indexed Indirect Type"
	case indirect_indexed_type:
		return "Indirect Indexed Type"
	default:
		return "Unknown Type"
	}
}

func printCodeType(ct code_type) string {
	switch ct {
	case function_definition:
		return "Function Definition"
	case op_code:
		return "Op Code"
	default:
		return "Unknown Type"
	}
}

func printOpCode(op Op_Code) string {
	switch op {
	case ADC_cmd:
		return "ADC"
	case AND_cmd:
		return "AND"
	case ASL_cmd:
		return "ASL"
	case BCC_cmd:
		return "BCC"
	case BCS_cmd:
		return "BCS"
	case BEQ_cmd:
		return "BEQ"
	case BIT_cmd:
		return "BIT"
	case BMI_cmd:
		return "BMI"
	case BNE_cmd:
		return "BNE"
	case BPL_cmd:
		return "BPL"
	case BRK_cmd:
		return "BRK"
	case BVC_cmd:
		return "BVC"
	case BVS_cmd:
		return "BVS"
	case CLC_cmd:
		return "CLC"
	case CLD_cmd:
		return "CLD"
	case CLI_cmd:
		return "CLI"
	case CLV_cmd:
		return "CLV"
	case CMP_cmd:
		return "CMP"
	case CPX_cmd:
		return "CPX"
	case CPY_cmd:
		return "CPY"
	case DEC_cmd:
		return "DEC"
	case DEX_cmd:
		return "DEX"
	case DEY_cmd:
		return "DEY"
	case EOR_cmd:
		return "EOR"
	case INC_cmd:
		return "INC"
	case INX_cmd:
		return "INX"
	case INY_cmd:
		return "INY"
	case JMP_cmd:
		return "JMP"
	case JSR_cmd:
		return "JSR"
	case LDA_cmd:
		return "LDA"
	case LDX_cmd:
		return "LDX"
	case LDY_cmd:
		return "LDY"
	case LSR_cmd:
		return "LSR"
	case NOP_cmd:
		return "NOP"
	case ORA_cmd:
		return "ORA"
	case PHA_cmd:
		return "PHA"
	case PHP_cmd:
		return "PHP"
	case PLA_cmd:
		return "PLA"
	case PLP_cmd:
		return "PLP"
	case ROL_cmd:
		return "ROL"
	case ROR_cmd:
		return "ROR"
	case RTI_cmd:
		return "RTI"
	case RTS_cmd:
		return "RTS"
	case SBC_cmd:
		return "SBC"
	case SEC_cmd:
		return "SEC"
	case SED_cmd:
		return "SED"
	case SEI_cmd:
		return "SEI"
	case STA_cmd:
		return "STA"
	case STX_cmd:
		return "STX"
	case STY_cmd:
		return "STY"
	case TAX_cmd:
		return "TAX"
	case TAY_cmd:
		return "TAY"
	case TSX_cmd:
		return "TSX"
	case TXA_cmd:
		return "TXA"
	case TXS_cmd:
		return "TXS"
	case TYA_cmd:
		return "TYA"
	default:
		return "Unknown Command"
	}
}

func dump_program() {
	var const_ptr *Const
	const_ptr = &const_head
	for const_ptr.next != nil {
		println(const_ptr.name, ": ", const_ptr.val)
		const_ptr = const_ptr.next
	}
	var ptr *Program_Code
	ptr = &head
	for ptr.next != nil {
		index := fmt.Sprintf("%x", ptr.index)
		if ptr.code_type == function_definition {
			fmt.Println(index, ": ", printCodeType(ptr.code_type), ptr.destination)
		} else {
			mem := fmt.Sprintf("%x", ptr.mem_1)
			fmt.Println(index, ": ", printCodeType(ptr.code_type), printOpCode(ptr.op_code), printAddrModeType(ptr.addr_mode), mem)
		}
		ptr = ptr.next
	}

}
