package library

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Homework900/simple-library2/domain"
)

type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	lastBookID   int
	lastReaderID int
}

func New() *Library {
	return &Library{
		Books:   make([]*domain.Book, 0),
		Readers: make([]*domain.Reader, 0),
	}
}

func (lib *Library) AddReader(firstName, lastName string) (*domain.Reader, error) {
	cleanFirstName := strings.TrimSpace(firstName)
	cleanLastName := strings.TrimSpace(lastName)
	if cleanFirstName == "" || cleanLastName == "" {
		return nil, errors.New("фамилия и имя не могут быть пустыми")
	}

	lib.lastReaderID++
	newReader := &domain.Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	lib.Readers = append(lib.Readers, newReader)
	return newReader, nil
}

func (lib *Library) AddBook(title, author string, year int) (*domain.Book, error) {
	cleanTitle := strings.TrimSpace(title)
	cleanAuthor := strings.TrimSpace(author)
	if cleanTitle == "" || cleanAuthor == "" {
		return nil, errors.New("название и автор не могут быть пустыми")
	}
	if year <= 0 {
		return nil, errors.New("год должен быть положительным числом")
	}

	lib.lastBookID++
	newBook := &domain.Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}
	lib.Books = append(lib.Books, newBook)
	return newBook, nil
}

func (lib *Library) DeleteBook(id int) error {
	for i, b := range lib.Books {
		if b.ID == id {
			if b.IsIssued {
				return fmt.Errorf("нельзя удалить книгу '%s' — она выдана читателю", b.Title)
			}
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена в библиотеке", id)
}

func (lib *Library) FindBookByTitle(title string) ([]*domain.Book, error) {
	if len(lib.Books) == 0 {
		return nil, fmt.Errorf("библиотека не содержит книг")
	}
	cleanedTitle := strings.ToLower(strings.TrimSpace(title))
	if cleanedTitle == "" {
		return nil, errors.New("название книги не может быть пустым")
	}
	result := []*domain.Book{}
	for _, book := range lib.Books {
		if book == nil {
			continue
		}
		if strings.ToLower(strings.TrimSpace(book.Title)) == cleanedTitle {
			result = append(result, book)
		}
	}
	return result, nil
}

func (lib *Library) FindBooksByAuthor(author string) ([]*domain.Book, error) {
	cleaned := strings.ToLower(strings.TrimSpace(author))
	if cleaned == "" {
		return nil, errors.New("автор не может быть пустым")
	}
	result := []*domain.Book{}
	for _, book := range lib.Books {
		if book == nil {
			continue
		}
		if strings.Contains(strings.ToLower(book.Author), cleaned) {
			result = append(result, book)
		}
	}
	return result, nil
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (lib *Library) GetAllReaders() []*domain.Reader {
	return lib.Readers
}

func (lib *Library) GetIssuedBooks() []*domain.Book {
	result := []*domain.Book{}
	for _, b := range lib.Books {
		if b.IsIssued {
			result = append(result, b)
		}
	}
	return result
}

func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}
	if err := book.IssueBook(reader); err != nil {
		return err
	}
	reader.AssignBook(book)
	return nil
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func (lib *Library) DeactivateReader(id int) error {
	reader, err := lib.FindReaderByID(id)
	if err != nil {
		return err
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s уже не активен", reader.FirstName, reader.LastName)
	}
	for _, b := range lib.Books {
		if b.IsIssued && b.ReaderID != nil && *b.ReaderID == id {
			return fmt.Errorf("нельзя деактивировать: у читателя на руках книга '%s'", b.Title)
		}
	}
	reader.Deactivate()
	return nil
}

func (lib *Library) ActivateReader(id int) error {
	reader, err := lib.FindReaderByID(id)
	if err != nil {
		return err
	}
	if reader.IsActive {
		return fmt.Errorf("читатель %s %s уже активен", reader.FirstName, reader.LastName)
	}
	reader.Activate()
	return nil
}

func (lib *Library) GetAllBooks() []*domain.Book {
	return lib.Books
}

type Stats struct {
	TotalBooks    int
	IssuedBooks   int
	AvailableBook int
	TotalReaders  int
	ActiveReaders int
	OverdueBooks  int
}

func (lib *Library) GetStats() Stats {
	s := Stats{
		TotalBooks:   len(lib.Books),
		TotalReaders: len(lib.Readers),
	}
	for _, b := range lib.Books {
		if b.IsIssued {
			s.IssuedBooks++
		} else {
			s.AvailableBook++
		}
	}
	for _, r := range lib.Readers {
		if r.IsActive {
			s.ActiveReaders++
		}
	}
	s.OverdueBooks = len(lib.GetOverdueBooks())
	return s
}

func (lib *Library) GetOverdueBooks() []*domain.Book {
	result := []*domain.Book{}
	for _, b := range lib.Books {
		if !b.IsIssued || b.ReaderID == nil {
			continue
		}
		reader, err := lib.FindReaderByID(*b.ReaderID)
		if err != nil {
			continue
		}
		if !reader.IsActive {
			result = append(result, b)
		}
	}
	return result
}

func (lib *Library) UpdateIDs() {
	for _, b := range lib.Books {
		if b.ID > lib.lastBookID {
			lib.lastBookID = b.ID
		}
	}
	for _, r := range lib.Readers {
		if r.ID > lib.lastReaderID {
			lib.lastReaderID = r.ID
		}
	}
}