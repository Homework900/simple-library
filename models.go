package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
}
func (b Book) String() string {
	return fmt.Sprintf("\"%s (%s, %d)\"", b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook(r *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
	if !r.IsActive {
		fmt.Printf("Читатель %s не активен и не может получить книгу.\n", r.Name)
		return
	}
	b.IsIssued = true
	fmt.Printf("Книга %s была выдана читателю %s\n", b.Title, r.Name)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

type Reader struct {
	ID       int
	Name     string
	IsActive bool
}