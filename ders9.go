package main

import "fmt"

func main() {
	/*	var x [5]int
		x[0] = 15

		fmt.Println(x)
	*/
	/*var x [5]float64

	x[0] = 42
	x[1] = 40
	x[2] = 30
	x[3] = 20
	x[4] = 10
	*/
	x := [5]float64{42, 40, 30, 20, 10}

	var toplam float64 = 0

	for i := 0; i < 5; i++ {

		toplam += x[i]
	}

	fmt.Println("Dizide ki elemanların toplamı", toplam)

	fmt.Println("Dizide ki elemanların ortalamaları=", toplam/float64(len(x)))
}
