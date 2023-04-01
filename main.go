package main

import "fmt"

func main() {
	// var sentence string
	// fmt.Println("Masukkan kalimat: ")
	sentence := "selamat malam"
	fmt.Scanln("%s", &sentence)

	for _, char := range sentence {
		fmt.Println(string(char))
	}

	charCount := make(map[string]int)
	for _, char := range sentence {
		charCount[string(char)]++
	}
	
	fmt.Println(charCount)
}