package main

import "fmt"

func main() {
	var d1 int = 42
	var d2 int = 6

	var toplamdeger int = d1 + d2
	var çıkarmadeger int = d1 - d2
	var çarpmadeger int = d1 * d2
	var bölmedeger int = d1 / d2

	fmt.Println("d1+d2", toplamdeger)
	fmt.Println("d1-d2", çıkarmadeger)
	fmt.Println("d1*d2", çarpmadeger)
	fmt.Println("d1/d2", bölmedeger)

}
