package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(`
	---------> Switch <----------
	Switch iboralari ko'plab tarmoqlar bo'ylab shartlarni ifodalaydi.
	`)
	number := 2
	switch number {
	case 1:
		fmt.Println("Bir")
	case 2:
		fmt.Println("ikki")
	case 3:
		fmt.Println("uch")
	}
	fmt.Println(time.Now().Year())
	switch time.Now().Year() {
	case 2022:
		fmt.Println("bu 2022")
	case 2025:
		fmt.Println("Yil &v ga ting", time.Now().Year())
	default:
		fmt.Println("Bilmadim")
	}
}
