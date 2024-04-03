package main

type Program_Code struct {
	code_type   int
	op_code     Op_Code
	index       int
	next        *Program_Code
	addr_mode   int
	mem_1       int
	destination string
}

type Op_Code int

type code_type int

type addr_mode int

const (
	implicit_type = iota
	accumulator_type
	immediate_type
	zero_page_type
	zero_page_x_type
	zero_page_y_type
	relative_type
	absolute_type
	absolute_x_type
	absolute_y_type
	indirect_type
	indexed_indirect_type
	indirect_indexed_type
)

const (
	function_start = iota
	op_code
)

const (
	ADC_cmd = iota
	AND_cmd
	ASL_cmd
	BCC_cmd
	BCS_cmd
	BEQ_cmd
	BIT_cmd
	BMI_cmd
	BNE_cmd
	BPL_cmd
	BRK_cmd
	BVC_cmd
	BVS_cmd
	CLC_cmd
	CLD_cmd
	CLI_cmd
	CLV_cmd
	CMP_cmd
	CPX_cmd
	CPY_cmd
	DEC_cmd
	DEX_cmd
	DEY_cmd
	EOR_cmd
	INC_cmd
	INX_cmd
	INY_cmd
	JMP_cmd
	JSR_cmd
	LDA_cmd
	LDX_cmd
	LDY_cmd
	LSR_cmd
	NOP_cmd
	ORA_cmd
	PHA_cmd
	PHP_cmd
	PLA_cmd
	PLP_cmd
	ROL_cmd
	ROR_cmd
	RTI_cmd
	RTS_cmd
	SBC_cmd
	SEC_cmd
	SED_cmd
	SEI_cmd
	STA_cmd
	STX_cmd
	STY_cmd
	TAX_cmd
	TAY_cmd
	TSX_cmd
	TXA_cmd
	TXS_cmd
	TYA_cmd
)
