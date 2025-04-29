package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Андрей"
	var age int = 25
	temperature := 36.6
	var onlineStatus bool = true
	cityName := "Екатеринбург"
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Температура:", temperature)
	fmt.Println("Статус:", onlineStatus)
	fmt.Println("Город:", cityName)
	fmt.Println(reflect.TypeOf(age))

}

