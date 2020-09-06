package core

// Core ...
type Core struct {
	Requests *Requests
}

// New ...
func New() *Core {
	return &Core{
		Requests: &Requests{},
	}
}
