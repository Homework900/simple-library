package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	id1 := 1
	reader2ID := 2

	user1 := Reader{
		ID:        &id1,
		FirstName: "Петр",
		LastName:  "Петров",
		IsActive:  true,
	}

	reader2 := Reader{
		ID:        &reader2ID,
		FirstName: "Sergey",
		LastName:  "Meniaylo",
		IsActive:  true,
	}

	book1 := Book{
		ID:     1,
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	}

	book1.IssueBook(&user1)
	fmt.Println(book1) 
	user1.AssignBook(&book1)

	book1.IssueBook(&reader2)

	user1.Deactivate()
	fmt.Println(user1)
	fmt.Println("---")
	book1.IssueBook(&user1)

	book1.ReturnBook()
	fmt.Println(book1)
	lib := Library{
		Books:   map[int]*Book{book1.ID: &book1},
		Readers: map[int]*Reader{*user1.ID: &user1, *reader2.ID: &reader2},
	}
	fmt.Printf("Книг в библиотеке: %d, читателей: %d\n",
		len(lib.Books), len(lib.Readers))
}