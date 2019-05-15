package module

// Module ...
type Module [11]*Section

// Section ...
type Section struct {
	ID       SectionID
	Size     SectionSize
	Contents []byte
}
