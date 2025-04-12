package main

import "fmt"

const (
	age    int    = 25
	name   string = "Asadbek"
	faolmi bool   = true
)

func main() {
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(faolmi)

	var ism string = "Asadbek"
	ism = "var katil so'zi orqali e'lon qilingan o'zgarovchilarni o'zgartirdim"
	fmt.Println(ism)

	const (
		a = 15
		b = 15
		c = 20
	)
	fmt.Println(a, b, c)

	const phione_number, age = "+998979437674", 23
	fmt.Println(phione_number, age)
}
