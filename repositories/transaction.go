package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetCurrentCarts(userID int) ([]models.Cart, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
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

// func (r *repository) UpdateTransaction(status string, ID string) error {
// 	var transaction models.Transaction
// 	r.db.First(&transaction, ID)

// 	// new status : pending
// 	// status : pending

// 	// If is different & Status is "success" decrement product quantity
// 	if status != transaction.Status && status == "success" {

// 		var order []models.Order
// 		orderedProduct, _ := r.GetTransactionProducts(order, transaction.ID)
// 		for _, p := range orderedProduct {
// 			var product models.Book
// 			r.db.First(&product, p.ID)
// 			r.db.Save(&product)
// 		}
// 	}

// 	transaction.Status = status

// 	err := r.db.Save(&transaction).Error

// 	return err
// }

func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Book").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var product models.Book
		r.db.First(&product, transaction.Book)
		r.db.Save(&product)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}
