package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("dairenin alanı =", dairealan(5))
	fmt.Println("dikdortgen alanı =", dikdörtgenalan(3, 2))

}

//dairenin alanını hesaplayan fonk

func dairealan(yariçap float64) float64 {

	return math.Pi * yariçap * yariçap
}

//dikdortgen alanı hesaplayan fonk
func dikdörtgenalan(a, b int) int {
	return a * b

}
