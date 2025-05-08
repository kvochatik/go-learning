package main

import "fmt"

type book struct {
	title string
	author string
	pages int
}

func main(){
	gore := book{"Горе от ума", "Пушкин", 450}
	more := book{
		title: "Гари поттер",
		author: "Телка",
		pages: 25,
	}
	
	fmt.Println(gore)
	fmt.Println(more)

	more.pages = 50

	fmt.Println(more)

	arr := []book{gore, more }
	for _, i := range arr{
		fmt.Println(i)
	}
}