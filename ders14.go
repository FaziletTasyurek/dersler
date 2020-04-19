package main

import "fmt"

//ana fonksiyon
func main() {

	//Hangi fonk önce çağrılırsa onu ilk önce yazacak
	defer ikinci() //defer bu fonksiyonu atla bir sonrakini çalıştır sonra buna geri dön anlamında kullanılır
	birinci()
	üçüncü()
}

func birinci() {

	fmt.Println("birinci fonksiyon")

}
func ikinci() {

	fmt.Println("ikinci fonksiyon")

}
func üçüncü() {

	fmt.Println("üçüncü fonksiyon")

}
