package main

import "fmt"

func Ex3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println("Ex3")
	fmt.Printf("Byte -> %v\nSmall Int -> %v\nBig uInt -> %v\n", b, smallI, bigI)
}
