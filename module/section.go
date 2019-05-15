package module

// SectionID ...
type SectionID byte

// SectionID values
const (
	CustomSection SectionID = iota
	TypeSection
	ImportSection
	FunctionSection
	TableSection
	MemorySection
	GlobalSection
	ExportSection
	StartSection
	ElementSection
	CodeSection
	DataSection
)

// SectionSize ...
type SectionSize uint32
