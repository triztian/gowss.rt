package instr

// Variable
// https://webassembly.github.io/spec/core/bikeshed/index.html#variable-instructions%E2%91%A0
const (
	LocalGet  OpCode = 0x20
	LocalSet         = 0x21
	LocalTee         = 0x22
	GlobalGet        = 0x32
	GlobalSet        = 0x24
)
