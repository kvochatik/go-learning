package main

import "fmt"

type book struct {
	title string
	author string
	pages int
}

func main(){
	more := book{
		title: "Гари поттер",
		author: "Телка",
		pages: 25,
	}
	fmt.Println(more)
	more.printInfo()
	more.editorPages(50)
	fmt.Println(more)
}

func (b *book) printInfo() {
	fmt.Printf("Книга называется %s ее автор %s и в ней %d страниц \n", b.title, b.author, b.pages)
}

func (b *book) editorPages(x int) {
	b.pages = x
	fmt.Printf("Теперь в книге %d страниц \n", x)
}