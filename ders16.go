//bir txt dosyadan içinde ki verileri okuma
//C:\Users\FAZİLET\Desktop\test.txt

package main

import (
	"fmt"
	"io/ioutil" //dosya işlaemlerimde kullanılır
)

func main() {
	path := "C:/Users/FAZİLET/Desktop/ders/deneme2" //dosyanın yolunu kopyalayıp yapıştırıyoruz

	bs, err := ioutil.ReadFile(path) //dosya okuma işlemleri

	if err != nil { //eğer error boş değilse
		fmt.Println("Dosya Bulunamadı")
		return
	}
	str := string(bs) //bir sıkıntı olamdan dosya okunduysa

	fmt.Println(str)

}
