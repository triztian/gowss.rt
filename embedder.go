package wasmrt

// Embedder is a "Low Level" interface that is defined by the following section
// in the Specification. Most users would want to use the `Runtime` type from 
// this same package.
// See:
//		https://webassembly.github.io/spec/core/bikeshed/index.html#a1-embedding
//
type Embedder interface {

	StoreInit() Store

	// Modules 

	ModuleDecode(...byte) (Module, error)

	ModuleParse(string) (Module, error)

	ModuleValidate(Module) error

	ModuleInstantiate(Store, Module, ...ExternalValue) (Store, ModuleInstance, error)

	ModuleImports(Module) []struct {
		ModuleName string
		ImportName string
		ExternType ExternalType
	}

	ModuleExports(Module) []struct {
		SymbolName string
		SymbolType ExternalType
	}

	InstanceExport(ModuleInstance, string) (ExternalValue, error)

	// Functions

	FuncAlloc(Store, FuncType, HostFunc)

	FuncType(Store, Addr) FuncType

	FuncInvoke(Store, Addr, ...Value) (Store, Value, error)

	// Tables

	TableAlloc(Store, TableType) (Store, Addr)

	TableType(Store, Addr) TableType

	TableRead(Store, Addr, uint32) (Addr, error)

	TableWrite(Store, Addr, uint32, Addr) (Store, error)

	TableSize(Store, Addr) uint32

	TableGrow(Store, Addr, uint32) (Store, error)

	// Memories 

	MemAlloc(Store, MemType) (Store, Addr)

	MemType(Store, Addr) MemType

	MemRead(Store, Addr, uint32) (byte, error)

	MemWrite(Store, Addr, uint32, byte) (Store, error)

	MemSize(Store, Addr) uint32

	// Globals 

	GlobalAlloc(Store, GlobalType, Value) (Store, Addr)

	GlobalType(Store, Addr) GlobalType

	GlobalRead(Store, Addr) Value

	GlobalWrite(Store, Addr, Value)  (Store, error)

}

type Addr uint64

// -- Stub types --

type Value interface{}

type Store interface{}

type Module interface{}

type ModuleInstance interface{}

type ExternalValue interface{}

type ExternalType interface{}

type FuncType interface{}

type HostFunc interface{}

type TableType interface{}

type MemType interface{}

type GlobalType interface{}