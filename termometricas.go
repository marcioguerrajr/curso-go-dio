package main

import "fmt"

const ebulicaoC float64 = 100.0

func main() {

	var C = ebulicaoC

	K := C + 273

	fmt.Println("Ponto de ebulição da água em é:", K, "°K")
	fmt.Println("Ponto de ebulição da água em é:", C, "°C")

	fmt.Printf("Ponto de ebulição da água em é %g°K ou %g°C", K, C)
}
