package interfaces

// Beverage é a interface que todas as bebidas devem implementar
type Beverage interface {
	Cost() float64
	Description() string
}
