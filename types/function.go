package types

// ValueType ...
type ValueType struct {
	I32 *i32
	I64 *i64
	F32 *f32
	F64 *f64
}

// Function ...
type Function struct {
	Params []ValueType
	Result [1]ValueType
}
