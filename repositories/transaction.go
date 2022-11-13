package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetCurrentCarts(userID int) ([]models.Cart, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransactionByID(ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Buyer").Preload("Book").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetCurrentCarts(userID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("User").Preload("Book").Preload("User").Where("user_id =?", userID).Find(&carts).Error

	return carts, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Create(&transaction).Error
	return transaction, err
}

func (r *repository) GetTransactionByID(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, ID).Error

	return transaction, err
}
