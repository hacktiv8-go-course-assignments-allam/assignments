package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama     string
	Alamat   string
	Pekerjaan string
}

func main() {
	teman1 := Teman{"Andi", "Jakarta", "Developer"}
	teman2 := Teman{"Budi", "Surabaya", "Designer"}
	teman3 := Teman{"Cindy", "Bandung", "Analyst"}
	teman4 := Teman{"Dini", "Yogyakarta", "Marketing"}

	temanList := map[int]Teman{
		1: teman1,
		2: teman2,
		3: teman3,
		4: teman4,
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run biodata.go <absen>")
		return
	}

	absen := args[0]
	var teman Teman
	var found bool

	for i, t := range temanList {
		if fmt.Sprintf("%d", i) == absen {
			teman = t
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Teman dengan absen %s tidak ditemukan\n", absen)
		return
	}

	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
}