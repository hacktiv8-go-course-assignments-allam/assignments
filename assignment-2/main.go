package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    // loop for i
    for i := 0; i < 5; i++ {
        fmt.Printf("nilai i  = %d\n", i)
    }

    // loop for j
    for j := 0; j < 11; j++ {
				if j == 5 {
					// array of characters
					// chars := []rune{'C', 'A', 'I', 'A', 'P', 'B', 'O'}
					chars := []rune{
						0x0421,
						0x0410,
						0x0428,
						0x0410,
						0x0420,
						0x0412,
						0x041E,
					}

					// map to store byte positions of characters
					charPos := make(map[rune]int)

					// print byte positions of characters in the order they appear in chars
					pos := 0
					for _, char := range chars {
							charPos[char] = pos
							pos += utf8.RuneLen(char)
							fmt.Printf("character U+%04X '%c' starts at byte position %d\n", char, char, charPos[char])
					}
				} else {
        	fmt.Printf("nilai j  = %d\n", j)
				}
    }
}