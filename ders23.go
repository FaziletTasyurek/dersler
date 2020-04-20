//pointer değerin hafızadaki yerini işaretliyor
package main

import "fmt"

func main() {

	a := 5  //a değişkenine 5 değerini atadık
	b := &a //b değişkenine anın adresini atadık

	fmt.Println("a değişkeninin değeri", a)          // a değerini yazdırdık
	fmt.Println("a değişkeninin hafızadaki yeri", b) //b değerini yazdırdık bu değer a nın yerini işarte eden değer
	fmt.Println("b değişkeninin değeri", *b)         //b nin göstermiş olduğu adresteki değeri yazdırdı

	//okuyup yazdırdığımız b ifadesini değiştirme şansımız da var

	//b değişkeninin yerini değiştirdiğimiz için otomatik olarak a ifadesini değitirmiş olacağız

	*b = 45
	fmt.Println("a değişkeninin değeri", a)

}
