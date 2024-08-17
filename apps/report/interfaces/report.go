package interfaces

// Report é a interface que todos os relatórios devem implementar.
type Report interface {
	Generate() string
	GetDescription() string
}
