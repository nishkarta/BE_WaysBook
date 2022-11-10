package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindBooks() ([]models.Book, error)
	GetBookByID(ID int) (models.Book, error)
	AddBook(book models.Book) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(book models.Book) (models.Book, error)
}

func RepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) GetBookByID(ID int) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, ID).Error

	return book, err
}

func (r *repository) AddBook(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) UpdateBook(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) DeleteBook(book models.Book) (models.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
