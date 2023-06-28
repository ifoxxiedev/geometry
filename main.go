package geometry

import "math"

const PI = 3.1416

/* Circ é responsável por calcular a área da circunferência */
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

/* Rect é responsável por calcular a área de um retangulo */
func Rect(base float64, altura float64) float64 {
	return base * altura
}

// Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
