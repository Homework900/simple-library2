package domain

import (
	"fmt"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int
}

func (b *Book) IssueBook(reader *Reader) error {
	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана", b.Title)
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен и не может получить книгу", reader.FirstName, reader.LastName)
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = nil
	return nil
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (r Reader) String() string {
	status := "активен"
	if !r.IsActive {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, № %d, %s", r.FirstName, r.LastName, r.ID, status)
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s\n", r.FirstName, r.LastName, b)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r *Reader) Activate() {
	r.IsActive = true
}