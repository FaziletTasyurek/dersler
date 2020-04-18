package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			fmt.Println("sıfır")
		case 1:
			fmt.Println("bir")
		}
	}
}
