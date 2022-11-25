package tools

type InterfaceManger interface {
	Init()
	IsConnected() bool
	ProvideDBConnection() any
}
