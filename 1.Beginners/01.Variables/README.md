# O'zgaruvchilar ma'lumotlar qiymatlarini saqlash uchun konteynerlardir

## Go'da o'zgaruvchini e'lon qilishning ikki yo'li mavjud

### 1. Var kalit so'zidan, keyin o'zgaruvchi nomi va turidan foydalaning

```go
package main

import ("fmt")

var ism = "Asadbek" // Birchi o'zgaruvchi

func main() {
 var meva = "Olma" // Ikkichi o'zgaruvchi
 fmt.Println(ism)
 fmt.Println(meva)
}
```

### 2. := belgisidan keyin oʻzgaruvchi qiymatidan foydalaning

```go
package main

import ("fmt")

func main() {
 varOne := 100
 varTwo := 2
 fmt.Println(varOne)
 fmt.Println(varTwo)
}
```

### Turlarni belgilash + Xulosa qilingan tur
Quyidagi kodda siz yozmoqchi bo‘lgan blok (grouped) o‘zgaruvchilar deklaratsiyasi Go tilida qanday yozilishi kerakligi va qanday imlo xatolar borligini ko‘rsatib tuzatib chiqaman:
<br>
```go
package main

import ("fmt")

func main() {
	var (
		son       int           // int tipidagi o'zgaruvchi, default qiymati 0
		raqam     int = 1       // int tipida va boshlang‘ich qiymat bilan
		salomlar  string = "salom" // string (qator) tipida qiymat berilgan
	)

	fmt.Println(son)
	fmt.Println(raqam)
	fmt.Println(salomlar)
}

```

### Ko'p o'zgaruvchilar deklaratsiyasiga o'ting

```go
package main
import ("fmt")

func main() {
 var one, two, three, four, five int = 1, 2, 3, 4, 5

 fmt.Println(bir)
 fmt.Println(ikki)
 fmt.Println(uch)
 fmt.Println(to'rtta)
}
```

### Blokdagi o'zgaruvchilar deklaratsiyasiga o'tish

```go
package main

import ("fmt")

func main() {
	var (
		num        int           // faqat deklaratsiya, default qiymati 0
		raqam      int = 1       // int tipida, qiymati 1
		salomlash  string = "salom" // string tipida, qiymati "salom"
	)

	fmt.Println(num)
	fmt.Println(raqam)
	fmt.Println(salomlash)
}
```

## Go o'zgaruvchilarni nomlash qoidalari

#### 1. Oʻzgaruvchi nomi harf yoki pastki chiziq belgisi (\_) bilan boshlanishi kerak.

#### 2. O‘zgaruvchi nomi raqam bilan boshlana olmaydi

#### 3. O‘zgaruvchi nomi faqat alfa-raqamli belgilar va pastki chiziqdan iborat bo‘lishi mumkin (a-z, A-Z, 0-9 va \_ )

#### 4. Oʻzgaruvchilar nomlari katta-kichik harflarga sezgir (yosh, Yosh va YOSH uch xil oʻzgaruvchidir)

#### 5. O‘zgaruvchi nomining uzunligiga cheklov yo‘q

#### 6. O‘zgaruvchi nomida bo‘shliq bo‘lishi mumkin emas

#### 7. O'zgaruvchi nomi hech qanday Go kalit so'zlari bo'lishi mumkin emas
