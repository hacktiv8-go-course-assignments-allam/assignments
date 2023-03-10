package main

import (
	"fmt"
	// "first-project/photo"
	// usr "first-project/user"
)

// file ini akan digunakan
// sebagai file utama yang akan
// dijalan kan oleh golang

// standard minimum suatu project
// 1. package main
// 2. function main

// constant: suatu tempat yang mengaliaskan dimana value di simpan dan tidak bisa berubah
// const const1 int = 100

// // const const2 uint64 = 100
// const pi float64 = 3.14

// // constant int yang akan auto
// // increment berdasarkan deklarasi awal
// const (
// 	cc1 = iota + 1 //1
// 	cc2            //2
// 	cc3            //3
// 	cc4            //4
// )

// const (
// 	cc5 = iota + 2 //2
// 	cc6            //3
// 	cc7            //4
// 	cc8            //5
// )

// PAYMENT TYPE
// const (
// 	VA            = iota + 1 //1
// 	CREDIT_CARD              //2
// 	BANK_TRANSFER            //3
// )


func assignment1Answer() {
	var i = 21
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	var j = true
	fmt.Println(j)
	fmt.Printf("%b\n", i)
	fmt.Println("?")
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("U+%X\n", 'Я')
	var k float64 = 123.456
	fmt.Printf("%.6f\n", k)
	fmt.Printf("%e\n", k)
	return
}

// hardening -> membentuk source code menjadi suatu binary file
func main() {
	assignment1Answer()
	//*****************************
	// comment !
	// fmt -> build in package yang biasa kita gunakan
	// untuk memunculkan out (ngeprint value)
	// ke terminal kita

	// syntax !
	// fmt.Println("Hello World from DTS Batch 7!") // akan ngeprint ke terminal dengan otomatis enter

	//*****************************
	// variable: suatu tempat yang mengaliaskan dimana value di simpan dan bisa berubah

	// basic
	// int, int8, int32, int64 -> angka bulat dan bisa negative
	// uint, uint8, uint32, uint64 -> angka bulat dan tidak bisa negative
	// float32, float64 -> bilangan koma dan bisa negative
	// string -> kata
	// bool -> true / false

	// advance
	// map
	// interface

	// ketika sudah jelas valuenya, bisa menggunakan :=
	// var1 := 100             // <- var1 adalah variable yang menampung value 100
	// var2 := "this is var 2" // <- var2 adalah variable yang menampung value bertipe string
	// var3 := uint64(100)

	// var1: int, var2: string
	// var4 := var1 + int(var3) // type cast -> mengubah tipe dari variable

	// ketika valuenya belum jelas, bisa menggunakan cara dibawah
	// var var5 float64
	// fmt.Printf("var1:%d var2:%s var3:%d var4:%d var5:%f\n", var1, var2, var3, var4, var5) // tidak otomatis enter

	//*****************************
	// var4 = var1 + const1 // var1:int const1:int -> valid
	// var6 := var3 + const2 + uint64(const1)

	// redeclare constant
	// const1 = 90 // tidak bisa dilakukan karena const1 adalah constanta

	// var5 = 100.7
	// fmt.Printf("var1:%v var2:%v var3:%v var4:%v var5:%f\n", var1, var2, var3, var4, var5)

	// luas lingkaran
	// circleArea := pi * var5 * var5
	// fmt.Println("luas area lingkaran: ", circleArea)

	// fmt.Println("test iota: ", cc1, cc2, cc3, cc4)
	// fmt.Println("test iota2: ", cc5, cc6, cc7, cc8)

	// variabel pembayaran
	// pay1 := 1
	// if pay1 == VA {
	// 	fmt.Println("pembayaran melalui Virtual Acc")
	// } else if pay1 == CREDIT_CARD {
	// 	fmt.Println("pembayaran melalui Credit Card")
	// } else if pay1 == BANK_TRANSFER {
	// 	fmt.Println("pembayaran melalui Bank Transfer")
	// }

	//*****************************
	// import internal package
	// user := 1
	// if user == usr.ADMIN {
	// 	fmt.Println("user adalah ADMIN user")
	// } else if user == usr.NORMAL {
	// 	fmt.Println("user adalah NORMAL user")
	// } else if user == usr.SUPER_ADMIN {
	// 	fmt.Println("user adalah SUPER ADMIN user")
	// }

	//*****************************
	// aritmatika
	// add1 := 100 + 1 // 101
	// modulo -> operator yang akan menghasilkan sisa bagi
	// mod1 := add1 % 100 // -> add1=101 dibagi 100 => 1, sisanya 1
	// fmt.Println("modulo 101 % 100=", mod1)

	// logical -> operator yang digunakan untuk membandingkan value
	// sehingga dia akan menghasilkan suatu boolean
	// buah1 := "apel"
	// buah2 := "jeruk"
	// buah3 := "semangka"

	// fmt.Println("bandingin buah1 == buah2: ", buah1 == buah2)                                     // false, buah1 apel buah2 jeruk, salah
	// fmt.Println("bandingin buah1 == buah2 && buah1 == buah3: ", buah1 == buah2 && buah1 == buah3) // false && false => false
	// fmt.Println("bandingin buah1 != buah2 || buah1 != buah3: ", buah1 != buah2 || buah1 != buah3) // true || true => true
	// fmt.Println("bandingin buah1 != buah2 || buah1 == buah3: ", buah1 != buah2 || buah1 == buah3) // true || false => true
	// fmt.Println("100 > 1:", 100 > 1)

	// bit operation
	// bit1 := 2
	// fmt.Println("bit wise >> 2", bit1>>2)
	// fmt.Println("bit wise << 2", bit1<<2)

	// MINI QUIZ
	// budi adalah user bertipe admin
	// dia memiliki foto sebanyak 10 foto
	// buatlah program yang akan menghasilkan output
	// 1. jika user adalah admin dan foto > 10 => akan print "ok"
	// 2. jika user adalah admin dan foto >= 10 => akan print "tidak ok"
	// 3. selain diatas => akan print "oke gak ya?"
	// kriteria:
	// type user disimpan di package user
	// batas foto disimpan di package photo
	// program akan dijalankan di main.go

	// budiType := 0 // admin 0
	// budiPhoto := 10

	// give white space
	// fmt.Println()

	// budiType == usr.ADMIN => true
	// budiPhoto > photo.BATAS_FOTO => false
	// budiPhoto >= photo.BATAS_FOTO => true

	// if budiType == usr.ADMIN && budiPhoto > photo.BATAS_FOTO {
	// 	fmt.Println("ok")
	// } else if budiType == usr.ADMIN && budiPhoto >= photo.BATAS_FOTO {
	// 	fmt.Println("tidak ok")
	// } else {
	// 	fmt.Println("oke gak ya?")
	// }
}

// func main() -> ketika di declarasi ulang,
// akan menimbulkan error

// func main1() -> ketika dijalankan, tidak akan
// terbaca sebagai main()