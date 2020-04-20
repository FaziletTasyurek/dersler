package main

import "fmt"

//struct yapısı tanımlanması

type personeller struct {
	adi, soyadi string
	yaş         int
	tamzamanlı  bool
	maaş        int
}

func main() {

	x := personeller{
		adi:        "fazilet",
		soyadi:     "taşyürek",
		yaş:        24,
		tamzamanlı: true,
		maaş:       2500,
	}
	fmt.Println(x)
	fmt.Println("adı", x.adi)
}
