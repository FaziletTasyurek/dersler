package main

import "fmt"

func main() {

	fmt.Println(sayitoplamlari(1, 2, 3, 4))
}

//fonksiyona birden fazla değer gönderip geriye değer toplamları döndüreceğiz

func sayitoplamlari(sayilar ...int) int { //fonksiyona birden fazla değer verebilmemizi sağlıyor
	toplam := 0
	for _, n := range sayilar {

		toplam += n
	}
	return toplam
}
