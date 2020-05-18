package vm

// Engine interface
type Engine interface {
	Execute() (interface{}, error)
}
