package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++	{
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		
		if i % 3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		
		if i % 5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}
}