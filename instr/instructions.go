package instr

// 32 bit integer instructions
const (
	// control
	Nop = iota
	Unreachable
	Block /* resulttype instr* end */
	loop  /* resulttype instr* end */
	If
	Else
	Br
	BrIf
	BrTable
	Return
	Call
	CallIndirect
	End

	// numeric

	// constants
	I32Const
	I64Const

	F32Const
	F64Const

	// Unary
	I32Clz
	I32Ctz
	I32Popcnt
	I64Clz
	I64Ctz
	I64Popcnt

	// Binary Integer
	I32Add
	I32Sub
	I32Mul
	I32DivUnsigned
	I32DivSigned
	I32And
	I32Or
	I32Xor
	I32Shl
	I32ShrUnsigned
	I32ShrSigned
	I32Rotl
	I32Rotr

	I64Add
	I64Sub
	I64Mul
	I64DivUnsigned
	I64DivSigned
	I64And
	I64Or
	I64Xor
	I64Shl
	I64ShrUnsigned
	I64ShrSigned
	I64Rotl
	I64Rotr

	// Unary Float
	F32Abs
	F32Neg
	F32Sqrt
	F32Ceil
	F32Floor
	F32Trunc
	F32Nearest

	F64Abs
	F64Neg
	F64Sqrt
	F64Ceil
	F64Floor
	F64Trunc
	F64Nearest

	// Binary Float
	F32Add
	F32Sub
	F32Mul
	F32Div
	F32Min
	F32Max
	F32CopySign

	F64Add
	F64Sub
	F64Mul
	F64Div
	F64Min
	F64Max
	F64CopySign

	// Tests
	I32Eqz
	I64Eqz

	I32Equal
	I32NotEqual
	I32LessThanSigned
	I32LessThanUnsigned
	I32GreaterThanSigned
	I32GreaterThanUnsigned
	I32LessEqualSigned
	I32LessEqualUnsigned
	I32GreaterEqualSigned
	I32GreaterEqualUnsigned
	I64Equal
	I64NotEqual
	I64LessThanSigned
	I64LessThanUnsigned
	I64GreaterThanSigned
	I64GreaterThanUnsigned
	I64LessEqualSigned
	I64LessEqualUnsigned
	I64GreaterEqualSigned
	I64GreaterEqualUnsigned

	F32Equal
	F32NotEqual
	F32GreaterThan
	F32LessThan
	F32LessEqual
	F32GreaterEqual

	I32WrapI64
	I64ExtendI32Signed
	I64ExtendI32Unsigned
	I32TruncateF32Signed
	I32TruncateF32Unsigned
	I32TruncateF64Signed
	I32TruncateF64Unsigned
	I64TruncateF32Signed
	I64TruncateF32Unsigned

	F32DemoteF64
	F64PromoteF32

	F32ConvertI32Signed
	F32ConvertI32Unsigned
	F32ConvertI64Signed
	F32ConvertI64Unsigned
	F64ConvertI32Signed
	F64ConvertI32Unsigned

	I32ReinterpretF64
	F32ReinterpretF64
	I64ReinterpretF32
	F64ReinterpretF32

	// Parametric
	// https://webassembly.github.io/spec/core/bikeshed/index.html#parametric-instructions%E2%91%A0
	Drop
	Select

	// Variable
	// https://webassembly.github.io/spec/core/bikeshed/index.html#variable-instructions%E2%91%A0
	LocalGet
	LocalSet
	LocalTee
	GlobalGet
	GlobalSet

	// Memory
	// https://webassembly.github.io/spec/core/bikeshed/index.html#memory-instructions%E2%91%A0
	I32Load
	F32Load
	I64Load
	F64Load

	I32Store
	F32Store
	I64Store
	F64Store

	I32Load8Signed
	I32Load8Unsigned
	I64Load8Signed
	I64Load8Unsigned

	I32Load16Signed
	I32Load16Unsigned

	I64Load16Signed
	I64Load16Unsigned

	I64Load32Signed
	I64Load32Unsigned

	I32Store8
	I64Store8

	I32Store16
	I64Store16

	I64Store32

	MemorySize
	MemoryGrow
)
