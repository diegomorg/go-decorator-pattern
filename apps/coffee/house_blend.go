package coffee

// Implementação concreta da interface Beverage
type HouseBlend struct{}

func (h *HouseBlend) Cost() float64 {
	return 0.89
}

func (h *HouseBlend) Description() string {
	return "House Blend Coffee"
}
