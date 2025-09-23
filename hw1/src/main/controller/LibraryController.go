package controller

import (
	"github.com/PetrosiunArtem/learning-golang/hw1/src/main/entity"
	"github.com/PetrosiunArtem/learning-golang/hw1/src/main/service"
)

type LibraryController interface {
	AddBookByTitle(book *entity.Book) (string, error)
	AddBookById(book *entity.Book) (string, error)
	GetBooksByKey(title string) (*entity.Book, error)
	GetAllBooks() []*entity.Book
}

type LibraryControllerImpl struct {
	libraryService service.LibraryService
}

func CreateLibraryController() LibraryController {
	return &LibraryControllerImpl{libraryService: service.CreateLibraryService()}
}

func (libraryController LibraryControllerImpl) AddBookByTitle(book *entity.Book) (string, error) {
	return libraryController.libraryService.AddBookByTitle(book)
}

func (libraryController LibraryControllerImpl) AddBookById(book *entity.Book) (string, error) {
	return libraryController.libraryService.AddBookById(book)
}
func (libraryController LibraryControllerImpl) GetBooksByKey(title string) (*entity.Book, error) {
	return libraryController.libraryService.GetBooksByKey(title)
}

func (libraryController LibraryControllerImpl) GetAllBooks() []*entity.Book {
	return libraryController.libraryService.GetAllBooks()
}
