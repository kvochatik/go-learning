package main

import "fmt"

func main() {
	fmt.Println(chet(4))//8
	defer fmt.Println(chet(6))// 12
	defer fmt.Println(chet(2))// 4
}

func chet (sum int) int {
	x := 2
	return x*sum
}