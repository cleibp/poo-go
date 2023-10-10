package main

import (
	"fmt"
	"math"
)

// Interface para representar uma forma
type Forma interface {
	CalcularArea() float64
}

// Estrutura para representar um retângulo
type Retangulo struct {
	Base   float64
	Altura float64
}

// Implementação do método CalcularArea para Retangulo
func (r Retangulo) CalcularArea() float64 {
	return r.Base * r.Altura
}

// Estrutura para representar um círculo
type Circulo struct {
	Raio float64
}

// Implementação do método CalcularArea para Circulo
func (c Circulo) CalcularArea() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	// Criar uma forma de retângulo
	retangulo := Retangulo{Base: 5.0, Altura: 4.0}

	// Criar uma forma de círculo
	circulo := Circulo{Raio: 3.0}

	// Calcular e imprimir a área do retângulo
	fmt.Printf("Área do Retângulo: %.2f\n", retangulo.CalcularArea())

	// Calcular e imprimir a área do círculo
	fmt.Printf("Área do Círculo: %.2f\n", circulo.CalcularArea())
}
