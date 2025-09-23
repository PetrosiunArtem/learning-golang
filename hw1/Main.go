package main

import (
	"fmt"

	"github.com/PetrosiunArtem/learning-golang/hw1/src/main/controller"
	"github.com/PetrosiunArtem/learning-golang/hw1/src/main/entity"
)

func main() {
	books := []*entity.Book{
		entity.NewBook("Book 1", "John Doe", 52),
		entity.NewBook("Book 2", "Luka", 42),
		entity.NewBook("Book 3", "Mike", 777),
	}
	LibraryController := controller.CreateLibraryController()
	for _, book := range books {
		_, err := LibraryController.AddBookByTitle(book)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	allBooks := LibraryController.GetAllBooks()
	for _, book := range allBooks {
		fmt.Print(book.Title, ", ")
	}
	fmt.Println()

	book1, _ := LibraryController.GetBooksByKey("book 1")
	book2, _ := LibraryController.GetBooksByKey("book 2")

	fmt.Printf("Title: %s, Author: %s, PageCount: %d \n", book1.Title, book1.Author, book1.PageCount)
	fmt.Printf("Title: %s, Author: %s, PageCount: %d \n", book2.Title, book2.Author, book2.PageCount)

	id3, _ := LibraryController.AddBookById(&entity.Book{Title: "Book 3", Author: "Artem", PageCount: 1440})
	id4, _ := LibraryController.AddBookById(&entity.Book{Title: "Book 4", Author: "Danil", PageCount: 1})

	book3, _ := LibraryController.GetBooksByKey(id3)
	book4, _ := LibraryController.GetBooksByKey(id4)

	fmt.Printf("Title: %s, Author: %s, PageCount: %d \n", book3.Title, book3.Author, book3.PageCount)
	fmt.Printf("Title: %s, Author: %s, PageCount: %d \n", book4.Title, book4.Author, book4.PageCount)
}
