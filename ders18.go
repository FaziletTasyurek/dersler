package main

import (
	"fmt"
	"time"
)

//bu gün değerini ay gün yı olarak ekrana yazdırma
func main() {
	d := time.Now()
	yıl, ay, gün := d.Date()

	fmt.Println("yıl =", yıl)
	fmt.Println("ay =", ay)
	fmt.Println("gün =", gün)

}
