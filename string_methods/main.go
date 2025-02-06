package main

import "fmt"

// Upper string
func Upper(str string) string {
	// Xotera- RAM sarfi Big-O(n)
	result := ""
	// CPU apperatsa Big-O(n)
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			ch = ch - 32
		}
		result += string(ch)
	}
	return result
}

// Lower string
func Lower(str string) string {
	result := ""
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			ch = ch + 32
		}
		result += string(ch)
	}
	return result
}

// is upper
func IsUpper(str string) bool {
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			return false
		}
	}
	return true
}

// is lower
func IsLower(str string) bool {
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("Hello")
	// upper string
	text := "Hello, Asadbek"
	fmt.Println(Upper(text))
	// lower string
	text_lower := "HELLO, Asadbek"
	fmt.Println(Lower(text_lower))
	// is upper
	fmt.Println(IsUpper(text))
	fmt.Println(IsLower(text_lower))
}
