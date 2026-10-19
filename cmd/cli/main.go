package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Homework900/simple-library2/library"
	"github.com/Homework900/simple-library2/storage"
)

const dbFile = "books.json"

func main() {
	myLibrary, err := storage.LoadLibraryFromJSON(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл данных не найден, создана новая библиотека.")
			myLibrary = library.New()
		} else {
			fmt.Println("Ошибка загрузки:", err)
			return
		}
	} else {
		fmt.Println("Библиотека успешно загружена из файла.")
	}

	runCLI(myLibrary)
}

func runCLI(lib *library.Library) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		scanner.Scan()
		choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Ошибка: введите число.")
			continue
		}

		if choice == 0 {
			fmt.Println("Сохранение данных и выход.")
			if err := storage.SaveLibraryToJSON(dbFile, lib); err != nil {
				fmt.Println("Ошибка сохранения:", err)
			} else {
				fmt.Println("Данные сохранены в", dbFile)
			}
			return
		}

		handleChoice(choice, lib, scanner)
	}
}

func printMenu() {
	fmt.Println("Главное меню:")
	fmt.Println("1. Поиск книги по названию")
	fmt.Println("2. Поиск книг по автору")
	fmt.Println("3. Выдать книгу")
	fmt.Println("4. Вернуть книгу")
	fmt.Println("5. Поиск читателя по номеру")
	fmt.Println("6. Показать список книг")
	fmt.Println("7. Показать книги на руках")
	fmt.Println("8. Показать всех читателей")
	fmt.Println("9. Добавить книгу")
	fmt.Println("10. Добавить читателя")
	fmt.Println("11. Удалить книгу по ID")
	fmt.Println("12. Статистика библиотеки")
	fmt.Println("13. Просроченные книги")
	fmt.Println("14. Деактивировать читателя")
	fmt.Println("15. Активировать читателя")
	fmt.Println("0. Выход")
	fmt.Println("Выберите пункт меню:")
}

func handleChoice(choice int, lib *library.Library, scanner *bufio.Scanner) {
	switch choice {

	case 1:
		fmt.Println("Введите название книги:")
		scanner.Scan()
		found, err := lib.FindBookByTitle(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка поиска:", err)
			return
		}
		if len(found) == 0 {
			fmt.Println("Ничего не найдено")
			return
		}
		for _, book := range found {
			fmt.Println(book)
		}

	case 2:
		fmt.Println("Введите автора (или часть):")
		scanner.Scan()
		found, err := lib.FindBooksByAuthor(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка поиска:", err)
			return
		}
		if len(found) == 0 {
			fmt.Println("Ничего не найдено")
			return
		}
		for _, book := range found {
			fmt.Println(book)
		}

	case 3:
		fmt.Println("Введите ID книги:")
		scanner.Scan()
		bookID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID книги должен быть числом")
			return
		}
		fmt.Println("Введите ID читателя:")
		scanner.Scan()
		readerID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID читателя должен быть числом")
			return
		}
		if err := lib.IssueBookToReader(bookID, readerID); err != nil {
			fmt.Println("Ошибка выдачи:", err)
			return
		}
		fmt.Printf("Книга с ID %d успешно выдана читателю с ID %d\n", bookID, readerID)

	case 4:
		fmt.Println("Введите ID книги:")
		scanner.Scan()
		bookID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID книги должен быть числом")
			return
		}
		if err := lib.ReturnBook(bookID); err != nil {
			fmt.Println("Ошибка возврата:", err)
			return
		}
		fmt.Printf("Книга с ID %d успешно возвращена\n", bookID)

	case 5:
		fmt.Println("Введите номер читателя:")
		scanner.Scan()
		readerID, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер читателя должен быть числом")
			return
		}
		reader, err := lib.FindReaderByID(readerID)
		if err != nil {
			fmt.Println("Ошибка поиска:", err)
			return
		}
		fmt.Println(reader)

	case 6:
		books := lib.GetAllBooks()
		if len(books) == 0 {
			fmt.Println("Библиотечный фонд пуст")
			return
		}
		fmt.Println("\n--- Список всех книг ---")
		for i, b := range books {
			fmt.Println(i+1, "\t", b)
		}

	case 7:
		issued := lib.GetIssuedBooks()
		if len(issued) == 0 {
			fmt.Println("Нет выданных книг")
			return
		}
		fmt.Println("\n--- Книги на руках ---")
		for _, b := range issued {
			fmt.Println(b)
		}

	case 8:
		readers := lib.GetAllReaders()
		if len(readers) == 0 {
			fmt.Println("Читателей нет")
			return
		}
		fmt.Println("\n--- Все читатели ---")
		for _, r := range readers {
			fmt.Println(r)
		}

 case 9:
        fmt.Println("Введите название книги:")
        scanner.Scan()
        title := scanner.Text()
        fmt.Println("Введите автора:")
        scanner.Scan()
        author := scanner.Text
        fmt.Println("Введите год издания:")
        scanner.Scan()
        year, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Год должен состоять из цифр")
            return
        }
        if _, err := lib.AddBook(title, author(), year); err != nil {
            fmt.Println("Ошибка добавления:", err)
            return
        }
        fmt.Println("Книга добавлена")

	case 10:
		fmt.Println("Введите имя:")
		scanner.Scan()
		firstName := scanner.Text()
		fmt.Println("Введите фамилию:")
		scanner.Scan()
		lastName := scanner.Text()
		if _, err := lib.AddReader(firstName, lastName); err != nil {
			fmt.Println("Ошибка регистрации:", err)
			return
		}
		fmt.Println("Читатель добавлен")

	case 11:
		fmt.Println("Введите ID книги для удаления:")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		if err := lib.DeleteBook(id); err != nil {
			fmt.Println("Ошибка удаления:", err)
			return
		}
		fmt.Println("Книга удалена")

	case 12:
		s := lib.GetStats()
		fmt.Println("\n--- Статистика ---")
		fmt.Printf("Всего книг: %d\n", s.TotalBooks)
		fmt.Printf("Выдано: %d\n", s.IssuedBooks)
		fmt.Printf("Доступно: %d\n", s.AvailableBook)
		fmt.Printf("Всего читателей: %d\n", s.TotalReaders)
		fmt.Printf("Активных читателей: %d\n", s.ActiveReaders)
		fmt.Printf("Просроченных книг: %d\n", s.OverdueBooks)

	case 13:
		overdue := lib.GetOverdueBooks()
		if len(overdue) == 0 {
			fmt.Println("Просроченных книг нет")
			return
		}
		fmt.Println("\n--- Просроченные книги ---")
		for _, b := range overdue {
			fmt.Println(b)
		}

	case 14:
		fmt.Println("Введите ID читателя для деактивации:")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		if err := lib.DeactivateReader(id); err != nil {
			fmt.Println("Ошибка деактивации:", err)
			return
		}
		fmt.Println("Читатель деактивирован")

	case 15:
		fmt.Println("Введите ID читателя для активации:")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		if err := lib.ActivateReader(id); err != nil {
			fmt.Println("Ошибка активации:", err)
			return
		}
		fmt.Println("Читатель активирован")

	default:
		fmt.Println("Неизвестный пункт меню")
	}
}