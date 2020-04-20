//goroutines ile channals(kanallar) aynı anlama gelmektedir tek fakları goroutines kavramında fonk birbirlerinden bağımsız olarak birbirlerine işlemler bitince sinyal göndermeden işlem bitirirken channels de olay tam tersi olarak işlenir.package dersler

package main

import "fmt"

//1)Kanal tanımlama
//2)kanala değer atama
//3)Kanaldan değer alma
func main() {

	/*	birincikanal :=make(chan string)//kanal tanımlama

		birincikanal <- "Merhaba Nasılsınız"//kanala değer atama

		değişken := <-birincikanal//kanaldaki değeri değişkene atadık
	*/

	s := []int{4, 6, 8, 7, 5, 2, 3}

	c := make(chan int)

	go topla(s[:3], c) //s dizisinin baştan 4 tanesini
	go topla(s[2:], c) //s dizisinin sondan 2 tanesi
	//2 kere bunu toplattıktan sonra toplanan bu iki değeri x ve y değişkenine atıyoruz
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
func topla(s []int, c chan int) {
	toplamdeğer := 0
	for i := 0; i < len(s); i++ {
		toplamdeğer += s[i]

	}
	c <- toplamdeğer //kanala değer atama
}
