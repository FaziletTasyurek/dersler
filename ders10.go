package main

import "fmt"

func main() {

	sayılar := []float64{42, 15, 30, 20, 70}
	fmt.Println("dizide ki değerlerin ortalamaları :", ortalama(sayılar))
}

func ortalama(x []float64) float64 {

	var toplam float64 = 0

	for i := 0; i < 5; i++ {
		toplam += x[i]

	}
	return toplam / float64(len(x))
}
