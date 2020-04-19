package main

import "fmt"

//closure fonksiyon örneği anonim olarak geçebiliyorlar
//fonk içinde fonk oluşturmayı sağlıyor

func main() {

	toplam := func(x, y int) int { //toplam burada bir değişken adı

		return x + y

	}
	fmt.Println(toplam(1, 3))

}
