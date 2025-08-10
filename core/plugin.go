package core

type Plugin interface {
	Name() string
	Middleware() Middleware
}
