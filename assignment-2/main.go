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
					chars := []rune{'C', 'A', 'I', 'A', 'P', 'B', 'O'}
					charMap := map[rune]rune{
						'C': 0x3DE,
						'A': 0x3CF,
						'I': 0x3DF,
						'P': 0x3D0,
						'B': 0x3D0,
						'O': 0x3CF,
					}


					// map to store byte positions of characters
					charPos := make(map[rune]int)

					// print byte positions of characters in the order they appear in chars
					pos := 0
					for _, char := range chars {
							charPos[char] = pos
							pos += utf8.RuneLen(char)

							// C => + 0x3DE
							// A => + 0x3CF
							// I => + 0x3DF
							// P => + 0x3D0
							// B => + 0x3D0
							// O => + 0x3CF
							fmt.Printf("character U+%04X '%c' starts at byte position %d\n", char + charMap[char], char, charPos[char] * 2)
					}
				} else {
        	fmt.Printf("nilai j  = %d\n", j)
				}
    }
}