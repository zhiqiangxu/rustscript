package vm

import (
	"fmt"

	"github.com/go-interpreter/wagon/exec"
	"github.com/go-interpreter/wagon/validate"
	"github.com/go-interpreter/wagon/wasm"
)

type engine struct {
	vm       *exec.VM
	wasmCode []byte
}

const (
	wasmMemLimit uint64 = 10 * 1024 * 1024
)

// NewEngine is ctor for Engine
func NewEngine(wasmCode []byte) Engine {
	return &engine{wasmCode: wasmCode}
}

const (
	methodName = "invoke"
)

func (e *engine) Execute() (ret interface{}, err error) {

	m, err := ReadWasmModule(e.wasmCode, newHostModule())
	if err != nil {
		return
	}

	err = validate.VerifyModule(m)
	if err != nil {
		return
	}

	entry, ok := m.Export.Entries[methodName]
	if !ok {
		err = fmt.Errorf("invoke not found")
		return
	}

	index := int64(entry.Index)

	vm, err := exec.NewVM(m)
	if err != nil {
		return
	}

	ret, err = vm.ExecCode(index)

	return
}

func newHostModule() *wasm.Module {
	return nil
}
