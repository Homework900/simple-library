package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int 
}

type Reader struct {
	ID        *int 
	Name      string
	FirstName string
	LastName  string
	IsActive  bool
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
}

func (b *Book) IssueBook(reader *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
	if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.\n",
			reader.FirstName, reader.LastName)
		return
	}
	b.ReaderID = reader.ID
	b.IsIssued = true
	fmt.Printf("Книга %s была выдана читателю %s %s\n",
		b.Title, reader.FirstName, reader.LastName)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
		return
	}
	b.ReaderID = nil
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

func (r Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s\n",
		r.FirstName, r.LastName, b.String())
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string {
	return fmt.Sprintf("Читатель %s %s (активен: %t)",
		r.FirstName, r.LastName, r.IsActive)
}

type Library struct {
	Books   map[int]*Book
	Readers map[int]*Reader
}