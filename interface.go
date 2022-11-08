package figuras

import "fmt"

type Figuras interface {
	Area() float64
	Perimetro() float64
}

func Medidas(f Figuras) {
	fmt.Println(f)
	fmt.Println(f.Area())
	fmt.Println(f.Perimetro())

}
