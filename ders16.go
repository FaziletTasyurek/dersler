//bir txt dosyadan içinde ki verileri okuma
//C:\Users\FAZİLET\Desktop\test.txt

package main

import (
	"fmt"
	"io/ioutil" //dosya işlaemlerimde kullanılır
)

func main() {
	path := "C:/Users/FAZİLET/Desktop/ders/denem.txt" //dosyanın yolunu kopyalayıp yapıştırıyoruz

	bs, err := ioutil.ReadFile(path) //dosya okuma işlemleri

	if err != nil {
		fmt.Println("Dosya Bulunamadı")
		return
	}
	str := string(bs)

	fmt.Println(str)

}
