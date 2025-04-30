package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9} // динамический массив

	for i := 0; i <= 10; i++ { // цикл выводит число от 1 до 10
		if i > 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i <=20; i++ {	//Выводим только четный
		if i%2 == 0 && i != 0{
			fmt.Println(i)
		}
	}

	for _, i := range arr{ // выводим числа из массива
		fmt.Println(i)
	}
}
