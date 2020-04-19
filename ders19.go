//sort paketi örneği
package main

import (
	"fmt"
	"sort"
)

func main() {
	//karışık şekilde atanmış harf dizinini sıralı bir şekilde ekrana yazdırma
	harfler := []string{"G", "H", "A", "K", "M"}

	sort.Strings(harfler)

	fmt.Println(harfler)

	//integer dizi içerisinde sıralama yapma
	sayılar := []int{3, 4, 5, 8, 0, 1, 2, 7}

	sort.Ints(sayılar)

	fmt.Println(sayılar)

	x := 4
	i := sort.Search(len(sayılar), func(i int) bool { return sayılar[i] >= x })
	if i < len(sayılar) && sayılar[i] == x {
		fmt.Printf("%d dizi içerisinde bulundu. Index No %d", x, i)
	} else {
		fmt.Printf("%d dizi içerisinde bulunamadı. ", x)
	}
}
