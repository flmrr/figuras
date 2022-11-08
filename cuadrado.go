package figuras

type Cuadrado struct {
	Lado float64
}

func (cuadrado *Cuadrado) Area() float64 {
	return cuadrado.Lado * cuadrado.Lado

}

func (cuadrado *Cuadrado) Perimetro() float64 {
	return 4 * cuadrado.Lado

}
