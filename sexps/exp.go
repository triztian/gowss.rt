package sexps

// Exp represents an S-Expression
type Exp struct {
	Atom interface{}
	A    *Exp
	B    *Exp
}
