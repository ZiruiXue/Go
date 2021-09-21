package main

import "fmt"

func main() {
	name := "bill"
	namePointer := &name
	fmt.Println(&namePointer) // namePointer address 0xc0000ac018
	// Pointers 和 slices 一样，pass by reference
	printPointer(namePointer) // 传递的是复制的 namePointer address 0xc0000ac028
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
