package main

import "fmt"

func main() {
	name := "Андрей"
	greet(name)
	fmt.Println(calc(20, 2))
}

func greet(name string) {
	fmt.Println("Привет,", name)
}

func calc(a, c int) (sum int, diff int, mult int, div float64) {
	sum = a + c
	diff = a - c
	mult = a * c
	div = float64(a) / float64(c)
	return
}
