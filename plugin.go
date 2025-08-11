package httpkit

type Plugin interface {
	Name() string
	Middleware() Middleware
}
