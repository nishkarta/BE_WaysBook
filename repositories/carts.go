package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	AddToCart(cart models.Cart) (models.Cart, error)
	GetPrice(bookID int) (book models.Book, err error)
	GetCartsByUser(userID int) ([]models.Cart, error)
	GetCartsByCurrentUser(userID int) ([]models.Cart, error)
	GetCartByID(ID int) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddToCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Preload("User").Preload("Books").Error
	return cart, err
}

func (r *repository) GetPrice(bookID int) (book models.Book, err error) {
	err = r.db.First(&book, bookID).Error
	return book, err
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("User").Preload("Book").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCartsByUser(userID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("User").Preload("Book").Where("user_id = ?", userID).Find(&carts).Error

	return carts, err
}
func (r *repository) GetCartsByCurrentUser(userID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("User").Preload("Book").Where("user_id = ?", userID).Find(&carts).Error

	return carts, err
}

func (r *repository) GetCartByID(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("User").Preload("Book").First(&cart, ID).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}
