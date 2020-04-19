package main

import (
	"errors"
	"fmt"
)

func main() {
	//Kendimize ait error nesnesi üretmek
	err := errors.New("Hatalı mesaj")
	fmt.Println(err)
}
