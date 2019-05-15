package embedder

// Embedder ...
type Embedder interface {
	StoreInit() Store

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

	FuncAlloc(Store, FuncType, HostFunc)
}

// -- Stub types --

type Store interface{}

type Module interface{}

type ModuleInstance interface{}

type ExternalValue interface{}

type ExternalType interface{}

type FuncType interface{}

type HostFunc interface{}
