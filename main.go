package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	b := Book{ID: 1, Title: "Война и мир", Author: "Толстой", Year: 1869}
	r := Reader{ID: 1, Name: "Иван", IsActive: true}

	fmt.Println(b)      // использует метод String()
	b.IssueBook(&r)     // выдаём книгу
	b.IssueBook(&r)     // уже выдана
	b.ReturnBook()      // возвращаем
	b.ReturnBook()      // уже в библиотеке
}