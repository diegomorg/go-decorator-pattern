package coffee

import (
	"diegomorg/go-decorator/apps/coffee/decorators"
	"diegomorg/go-decorator/apps/coffee/interfaces"
	"fmt"
)

func Run() {
	var beverage interfaces.Beverage

	// Cria uma bebida básica
	beverage = &Espresso{}
	log(beverage)

	// adiciona Mocha ao Espresso
	beverage = &decorators.Mocha{Decorator: decorators.Decorator{Beverage: beverage}}
	log(beverage)

	// adiciona Milk ao Espresso com Mocha
	beverage = &decorators.Milk{Decorator: decorators.Decorator{Beverage: beverage}}
	log(beverage)
}

func log(beverage interfaces.Beverage) {
	fmt.Printf("%s $%.2f\n", beverage.Description(), beverage.Cost())
}
