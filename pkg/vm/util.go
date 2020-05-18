package vm

import (
	"bytes"
	"fmt"

	"github.com/go-interpreter/wagon/wasm"
)

// ReadWasmModule ...
func ReadWasmModule(code []byte, env *wasm.Module) (m *wasm.Module, err error) {
	m, err = wasm.ReadModule(bytes.NewReader(code), func(name string) (em *wasm.Module, err error) {
		switch name {
		case "env":
			em = env
			return
		}
		err = fmt.Errorf("unknown module: %q", name)
		return
	})

	return
}
