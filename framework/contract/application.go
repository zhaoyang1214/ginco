package contract

type Application interface {
	Container
	Version() string
	BasePath(path string) string
	RuntimePath() string
	GetIgnore(name string) interface{}
}
