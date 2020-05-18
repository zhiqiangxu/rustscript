package vm

import (
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

func (e *engine) Execute() (ret interface{}, err error) {

	m, err := ReadWasmModule(e.wasmCode, newHostModule())
	if err != nil {
		return
	}

	err = validate.VerifyModule(m)
	if err != nil {
		return
	}

	vm, err := exec.NewVM(m)
	if err != nil {
		return
	}

	index := int64(m.Start.Index)
	ret, err = vm.ExecCode(index)

	return
}

func newHostModule() *wasm.Module {
	return nil
}
