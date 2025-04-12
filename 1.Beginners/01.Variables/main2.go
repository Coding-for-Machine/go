package main

import "fmt"

func main(){
	/*Go tilida o‘zgaruvchiga tur (type) biriktirish va kompilyator tomonidan avtomatik turni aniqlash (type inference) — bu tilning muhim xususiyatlaridan biri.*/
	/*🟢 1. Assigning Types (Tur biriktirish)
Bu usulda siz o‘zgaruvchi e’lon qilayotganda aniq tipini ko‘rsatasiz.*/
	
	// Misol:
	var ism string = "Asadbek"
	var yosh int = 25
	var faolmi bool = true
	fmt.Println(ism, yosh, faolmi)

	/*
	🧠 Afzalliklari:

Kod aniq va tushunarli bo‘ladi

Katta loyihalarda qat’iy tiplar yordam beradi

Boshqa dasturchilar uchun osonroq o‘qiladi
*/

  /*🟡 2. Type Inferred (Tip avtomatik aniqlanadi)
Bu usulda siz := operatoridan foydalanasiz va Go qiymatga qarab o‘zi tipni aniqlaydi.*/
	// Mison:
	age := 25         // int deb tushuniladi
        name := "Asadbek"   // string deb tushuniladi
        is_active := true    // bool deb tushuniladi
	fmt.Println(age, name, is_active)
     /*🔴 Eslatma:
:= faqat funksiya ichida ishlatiladi. Tashqarida faqat var ishlaydi.*/
}
