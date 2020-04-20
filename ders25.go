// range kavramı -> üzerinde kullandığı diziyi for döngüsü ile tekrarlamayı sağlıyor
//bir dil range edildiğinde tekrarlama başına iki değer döndürme imkanına sahip
//birinci değer dizide ki elemanın indeksi ikinci değer dizide ki elemanun içeriği olarak belirtebilir

package main

import "fmt"

func main() {
	var isimler = []string{"fazi", "fazilet", "faziş"} //dizi tanımladık

	for a, b := range isimler { //for dizisi ile isimler dizisinde dönerken range ile a değişkenine indeksini b değişkenine ise değerini atadık
		//range bize iki değer döndürmüş oldu

		fmt.Println(a,". index ", b)
	}
}
