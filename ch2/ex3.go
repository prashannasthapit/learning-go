package main

import "fmt"

func Ex3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Printf("Ex3 Byte -> %v Small Int -> %v Big uInt -> %v\n", b, smallI, bigI)
}
