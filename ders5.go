package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		if i%2 == 0 {

			fmt.Println(i, "2 ye tam böldü")

		} else if i%3 == 0 {

			fmt.Println(i, "3 e tam böldü")
		} else {
			fmt.Println(i, "diğer")
		}

		/*else {
			fmt.Println(i, "tek sayi")
		}*/

	}
}
