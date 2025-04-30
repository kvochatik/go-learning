package main

import "fmt"

func main(){
	number := 10
	number1 := 120
	number3 := -5
	fmt.Println("Число", number, chislo(number), "и", chet(number))
	fmt.Println("Число", number1, chislo(number1), "и", chet(number1))
	fmt.Println("Число", number3, chislo(number3), "и", chet(number3))
	
}

func chislo (number int) string {
	if number > 0 {
		return "Число положительное"
	} else if number < 0 {
		return "Число отрицательное"
	} else {
		return "Это ноль"
	}
}

func chet (number int) string {
	if number > 100 && number%2 == 0 {
		return "Большое число"
	} else if number < 100 || number%2 != 0 {
		return "Или маленькое, или нечетное"
	} 
	return "не прошло"
}