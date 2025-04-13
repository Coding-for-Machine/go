package main

import "fmt"

func main() {
	tushincha := `
	go da for loop - bu aylanma harakni bildiradi va boshlangich nuqta (start:=1), 
	tugash nuqtalari (end<number or end> number) va qadamlar yani sakrashlari bilgilash 
	shar chunki for loop chiksizlikka tushib qoladi.
	{ coding fo Machine }
	`
	fmt.Println(tushincha)
	// Misol 1:
	fmt.Println("Misol: 1")
	start := 1
	end := 10
	setup := 1
	for start < end {
		fmt.Println(start) // Natija : 1, 2, 3, 4, 5, 6, 7, 8, 9
		start += setup
	}
	//  misol 2:
	fmt.Println("Misol: 2")
	for i := 1; i < 10; i++ {
		fmt.Println("Salom", i)
	}
	// misol: 3
	fmt.Println("Misol: 3")
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++ // qadamlar esdan chimasin i+=2
	}
	fmt.Println(`Misol: 4
	0 dan 100 gacha 
	Juft Sonlarni Ekranga chiqarish.

	`)
	for i := 0; i <= 100; i += 2 {
		fmt.Println(i)
	}
	// misol 5; Range
	fmt.Println(`
	misol 5; Range

	`)
	for i := range 3 {
		fmt.Println(i)
	}
	// misol 6
	fmt.Println(`
	Misol: 6
	range orqali 10 gacha 
	bulgan jusf sonlarni chiqarish
	
	`)
	for i := range 10 {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
	// Misol: 7, break- bu tuxtatish yani tormiz
	fmt.Println(`
	Misol: 7 
	break - bu tormiz yani tuxtatish
	`)
	for i := 1; i < 10; i++ {
		if i == 8 {
			break
		}
		fmt.Println("Break 8 da", i)
	}

}
