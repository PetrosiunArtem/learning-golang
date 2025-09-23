package repository

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PetrosiunArtem/learning-golang/hw1/src/main/entity"
)

type LibraryRepository interface {
	AddBookByTitle(book *entity.Book) (string, error)
	AddBookById(book *entity.Book) (string, error)
	GetBooksByKey(title string) (*entity.Book, error)
	GetAllBooks() []*entity.Book
}
type LibraryRepositoryImpl struct {
	books map[string]*entity.Book
	count int
}

func CreateLibraryRepository() LibraryRepository {
	return &LibraryRepositoryImpl{books: make(map[string]*entity.Book), count: 0}
}

func (libraryRepository LibraryRepositoryImpl) AddBookByTitle(book *entity.Book) (string, error) {
	key := generateKeyV1(book.Title)
	if libraryRepository.books[key] != nil {
		return "", errors.New("book already exists")
	}
	libraryRepository.books[key] = book
	return key, nil
}

func (libraryRepository *LibraryRepositoryImpl) AddBookById(book *entity.Book) (string, error) {
	key := libraryRepository.generateKeyV2()
	if libraryRepository.books[key] != nil {
		return "", errors.New("book already exists")
	}
	libraryRepository.books[key] = book
	return key, nil
}

func (libraryRepository LibraryRepositoryImpl) GetBooksByKey(title string) (*entity.Book, error) {
	if book, exists := libraryRepository.books[title]; exists {
		return book, nil
	}
	return nil, errors.New("book not found by title: " + title)
}

func (libraryRepository LibraryRepositoryImpl) GetAllBooks() []*entity.Book {
	result := make([]*entity.Book, 0, len(libraryRepository.books))
	for _, book := range libraryRepository.books {
		result = append(result, book)
	}
	return result
}

func generateKeyV1(title string) string {
	return strings.ToLower(title)
}

func (libraryRepository *LibraryRepositoryImpl) generateKeyV2() string {
	key := libraryRepository.count
	libraryRepository.count++
	return strconv.Itoa(key)
}
