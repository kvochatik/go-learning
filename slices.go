package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:3]
	frut := []string{"яблоко", "груша", "неведанный фрукт"}
	frut = append(frut, "апельсин")
	l, c := len(arr), cap(arr)

	fmt.Println("Массив:", arr)
	fmt.Println("Слайс:", slice)
	fmt.Println("Фрукты:", frut)
	frut = append([]string{"Ярбуз"}, frut...)
	fmt.Println("Добавление фрукта в начало", frut)
	fmt.Printf("Длина массива: %d, Кописити: %d\n", l, c)

}
