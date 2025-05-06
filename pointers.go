package main

import "fmt"

func main(){
	x := 5
	p := &x
	fmt.Println("Адрес переменной: ", p)
	fmt.Println("Значение по указателю: ", *p)
	*p *= 2
	fmt.Println("Новое значение: ", x)
}