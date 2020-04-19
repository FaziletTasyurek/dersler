package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Fazilet Taşyürek", "Taşyürek")) //kelime içerisinde taşyürek kelimesi var mı kontrolü varsa true sonuç yoksa false sonuç döner
	fmt.Println(strings.Count("Fazilet Taşyürek", "a"))           //kelime içerisinde a kaçtane var kontrolü yapar
	fmt.Println(strings.Index("Fazilet Taşyürek", "a"))           //kelime içerisinde yer alan a kaçıncı index de yer alır

	//strings. diyip tek tek ne işlem yapıldığı öğrennilebilir

}
