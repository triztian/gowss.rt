package instr

// Memory
// https://webassembly.github.io/spec/core/bikeshed/index.html#memory-instructions%E2%91%A0
const (
	I32Load           OpCode = 0x28
	F32Load                  = 0x29
	I64Load                  = 0x2A
	F64Load                  = 0x2B
	I32Load8Signed           = 0x2C
	I32Load8Unsigned         = 0x2D
	I32Load16Signed          = 0x2E
	I32Load16Unsigned        = 0x2F

	I64Load8Signed    = 0x30
	I64Load8Unsigned  = 0x31
	I64Load16Signed   = 0x32
	I64Load16Unsigned = 0x33
	I64Load32Signed   = 0x34
	I64Load32Unsigned = 0x35

	I32Store = 0x36
	I64Store = 0x37
	F32Store = 0x38
	F64Store = 0x39

	I32Store8  = 0x3A
	I32Store16 = 0x3B
	I64Store8  = 0x3C
	I64Store16 = 0x3D
	I64Store32 = 0x3E

	MemorySize = 0x3F
	MemoryGrow = 0x40
)
