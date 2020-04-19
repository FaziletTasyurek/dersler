package main

import "fmt"

//ana fonksiyon
func main() {

	fmt.Println(faktoriyel(5))

}

//recursion fonk
func faktoriyel(a uint) uint {
	if a == 0 {
		return 1
	} else {
		return a * faktoriyel(a-1)
	}

}
