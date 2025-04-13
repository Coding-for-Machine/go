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
	// misollar .

	if 7%2==0 {
		fmt.Println("7 ti 2 ga bulinmaydi")
	}else {
		fmt.Println("7 ga 2 ga bulinadi")
	}
	if 8%4 {
		fmt.Println("8 bu son 4 ga bulinadi")
	}else {
		fmt.Println("8 bu son 4 ga bulinmaydi")
	}
	if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("8 yoki 7 juft son")
    	}

}
