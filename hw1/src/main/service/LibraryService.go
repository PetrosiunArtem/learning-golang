package service

import (
	"github.com/PetrosiunArtem/learning-golang/hw1/src/main/entity"
	"github.com/PetrosiunArtem/learning-golang/hw1/src/main/repository"
)

type LibraryService interface {
	AddBookByTitle(book *entity.Book) (string, error)
	AddBookById(book *entity.Book) (string, error)
	GetBooksByKey(title string) (*entity.Book, error)
	GetAllBooks() []*entity.Book
}

type LibraryServiceImpl struct {
	libraryRepository repository.LibraryRepository
}

func CreateLibraryService() LibraryService {
	return &LibraryServiceImpl{libraryRepository: repository.CreateLibraryRepository()}
}

func (libraryService LibraryServiceImpl) AddBookByTitle(book *entity.Book) (string, error) {
	key, err := libraryService.libraryRepository.AddBookByTitle(book)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (libraryService LibraryServiceImpl) AddBookById(book *entity.Book) (string, error) {
	key, err := libraryService.libraryRepository.AddBookById(book)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (libraryService LibraryServiceImpl) GetBooksByKey(title string) (*entity.Book, error) {
	book, err := libraryService.libraryRepository.GetBooksByKey(title)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (libraryService LibraryServiceImpl) GetAllBooks() []*entity.Book {
	return libraryService.libraryRepository.GetAllBooks()
}
