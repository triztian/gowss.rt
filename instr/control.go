package instr

// 32 bit integer instructions
const (
	// control
	Unreachable  OpCode = 0x00
	Nop                 = 0x01
	Block               = 0x02 /* resulttype instr* end */
	loop                = 0x03 /* resulttype instr* end */
	If                  = 0x04
	Else                = 0x05
	Br                  = 0x0C
	BrIf                = 0x0D
	BrTable             = 0x0E
	Return              = 0x0F
	Call                = 0x10
	CallIndirect        = 0x11
	End                 = 0x0B
)
