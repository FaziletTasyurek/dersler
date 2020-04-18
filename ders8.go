package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("bugün pazartesi")
	case time.Tuesday:
		fmt.Println("bugün salı")
	case time.Wednesday:
		fmt.Println("bugün çarşamba")
	case time.Thursday:
		fmt.Println("bugün perşembe")
	case time.Friday:
		fmt.Println("bugün cuma")
	case time.Saturday:
		fmt.Println("bugün cumartesi")
		//	case time.Sunday:fmt.Println("bugün pazar")
	default:
		fmt.Println("bugün pazar")
	}
}
