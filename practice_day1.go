package main

import "fmt"

func main() {
	fmt.Println(chet(4))
}

func chet (sum int) int {
	x := 2
	return x*sum
}