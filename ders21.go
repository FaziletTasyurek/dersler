//kod içerisnde yapılan görevlendirmeyi eş zamanlı olarak gerçekleştirme -->Goroutines

package main

import (
	"fmt"
	"time"
)

//bir değer beliryeceğiz belirlediğimiz değere kadar 0 dan başlayarak ekrana birer saniye aralıklarla yazdıracak
func yaz(d int) {
	for i := 0; i <= d; i++ {
		time.Sleep(time.Second) //ekranda bir saniye bekleme sağlayacak
		fmt.Println(i)
	}
}
func main() {
	//bu iki işlemi yapması 10sn fazla sürde ancal gorountines ile daha hızlı yapılabilir
	go yaz(5)
	yaz(5)
}
