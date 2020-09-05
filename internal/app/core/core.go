package core

// Core ...
type Core struct {
	Req *Requests
}

// New ...
func New() *Core {
	return &Core{
		Req: &Requests{},
	}
}
