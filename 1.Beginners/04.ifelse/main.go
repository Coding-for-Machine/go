package main

import "fmt"

func main() {
	fmt.Println("Sartlarimizni quyamiz dasturimizda!")

	sum := 6400
	if sum > 6400 {
		fmt.Println(sum, "Demaok boyman, 😁")
	} else {
		fmt.Println("Demak qanbag'alman")
	}
	// passwordni validatsa qilamiz
	password := "12551552414"

	if len(password) < 8 {
		fmt.Println("Password natog'ri!")
	} else {
		fmt.Println("Password tugri?")
	}
	//  Yoshni Validatsi qilamiz
	age := 23
	if age >= 18 {
		fmt.Println("Salom, siz tizimga kiraolasiz")
	} else {
		fmt.Println("Siz yossiz?")
	}
}
