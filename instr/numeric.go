package instr

// Numeric instructions
const (
	// constants
	I32Const OpCode = 0x41
	I64Const        = 0x42
	F32Const        = 0x43
	F64Const        = 0x44

	// Tests
	I32Eqz                  = 0x45
	I32Equal                = 0x46
	I32NotEqual             = 0x47
	I32LessThanSigned       = 0x48
	I32LessThanUnsigned     = 0x49
	I32GreaterThanSigned    = 0x4A
	I32GreaterThanUnsigned  = 0x4B
	I32LessEqualSigned      = 0x4C
	I32LessEqualUnsigned    = 0x4D
	I32GreaterEqualSigned   = 0x4E
	I32GreaterEqualUnsigned = 0x4F

	I64Eqz                  = 0x50
	I64Equal                = 0x51
	I64NotEqual             = 0x52
	I64LessThanSigned       = 0x53
	I64LessThanUnsigned     = 0x54
	I64GreaterThanSigned    = 0x55
	I64GreaterThanUnsigned  = 0x56
	I64LessEqualSigned      = 0x57
	I64LessEqualUnsigned    = 0x58
	I64GreaterEqualSigned   = 0x59
	I64GreaterEqualUnsigned = 0x5A

	F32Equal        = 0x5B
	F32NotEqual     = 0x5C
	F32GreaterThan  = 0x5D
	F32LessThan     = 0x5E
	F32LessEqual    = 0x5F
	F32GreaterEqual = 0x60

	F64Equal        = 0x61
	F64NotEqual     = 0x62
	F64GreaterThan  = 0x63
	F64LessThan     = 0x64
	F64LessEqual    = 0x64
	F64GreaterEqual = 0x66

	// Unary
	I32Clz    = 0x67
	I32Ctz    = 0x68
	I32Popcnt = 0x69

	// Binary Integer
	I32Add               = 0x6A
	I32Sub               = 0x6B
	I32Mul               = 0x6C
	I32DivSigned         = 0x6D
	I32DivUnsigned       = 0x6E
	I32RemainderSigned   = 0x6F
	I32RemainderUnsigned = 0x70
	I32And               = 0x71
	I32Or                = 0x72
	I32Xor               = 0x73
	I32Shl               = 0x74
	I32ShrSigned         = 0x75
	I32ShrUnsigned       = 0x76
	I32Rotl              = 0x77
	I32Rotr              = 0x78

	I64Clz               = 0x79
	I64Ctz               = 0x7A
	I64Popcnt            = 0x7B
	I64Add               = 0x7C
	I64Sub               = 0x7D
	I64Mul               = 0x7E
	I64DivSigned         = 0x7F
	I64DivUnsigned       = 0x80
	I64RemainderSigned   = 0x81
	I64RemainderUnsigned = 0x82
	I64And               = 0x83
	I64Or                = 0x84
	I64Xor               = 0x85
	I64Shl               = 0x86
	I64ShrSigned         = 0x87
	I64ShrUnsigned       = 0x88
	I64Rotl              = 0x89
	I64Rotr              = 0x8A

	// Unary Float
	F32Abs     = 0x8B
	F32Neg     = 0x8C
	F32Ceil    = 0x8D
	F32Floor   = 0x8E
	F32Trunc   = 0x8F
	F32Nearest = 0x90
	F32Sqrt    = 0x91

	// Binary Float
	F32Add      = 0x92
	F32Sub      = 0x93
	F32Mul      = 0x94
	F32Div      = 0x95
	F32Min      = 0x96
	F32Max      = 0x97
	F32CopySign = 0x98

	F64Abs     = 0x99
	F64Neg     = 0x9A
	F64Ceil    = 0x9B
	F64Floor   = 0x9C
	F64Trunc   = 0x9D
	F64Nearest = 0x9E
	F64Sqrt    = 0x9F

	F64Add      = 0xA0
	F64Sub      = 0xA1
	F64Mul      = 0xA2
	F64Div      = 0xA3
	F64Min      = 0xA4
	F64Max      = 0xA5
	F64CopySign = 0xA6

	I32WrapI64             = 0xA7
	I32TruncateF32Signed   = 0xA8
	I32TruncateF32Unsigned = 0xA9
	I32TruncateF64Signed   = 0xAA
	I32TruncateF64Unsigned = 0xAB
	I64ExtendI32Signed     = 0xAC
	I64ExtendI32Unsigned   = 0xAD
	I64TruncateF32Signed   = 0xAE
	I64TruncateF32Unsigned = 0xAF
	I64TruncateF64Signed   = 0xB0
	I64TruncateF64Unsigned = 0xB1

	F32ConvertI32Signed   = 0xB2
	F32ConvertI32Unsigned = 0xB3
	F32ConvertI64Signed   = 0xB4
	F32ConvertI64Unsigned = 0xB5

	F32DemoteF64 = 0xB6

	F64ConvertI32Signed   = 0xB7
	F64ConvertI32Unsigned = 0xB8
	F64ConvertI64Signed   = 0xB9
	F64ConvertI64Unsigned = 0xBA

	F64PromoteF32 = 0xBB

	I32ReinterpretF32 = 0xBC
	I64ReinterpretF64 = 0xBD
	F32ReinterpretI32 = 0xBE
	F64ReinterpretF64 = 0xBF
)
