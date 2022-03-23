package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//3 dasar pemrograman
	//1. program berjalan secara berurutan ( sequence)
	//2. terdapat percabangan (branching) ->IF
	//3. terdapat kondisi perulangan (looping) ->FOR

	// var namaDepan string = "Eric"
	// var namaBelakang string = "Widhi Antara"
	var umur int8 = 23

	if umur >= 18 {
		fmt.Println(umur, "Menuju dewasa")
	} else {
		fmt.Println(umur, "Masih bocah")
	}
}

//create program to display hello world
