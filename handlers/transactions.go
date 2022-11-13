package handlers

import (
	"encoding/json"
	"net/http"
	dto "waysbook/dto/results"
	transactiondto "waysbook/dto/transactions"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.FindTransactions()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: transactions}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) AddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	carts, _ := h.TransactionRepository.GetCurrentCarts(userId)

	cartLength := len(carts)
	total := 0
	// var bookIds []int

	for i := 0; i < cartLength; i++ {
		total += carts[i].Price
		// bookIds = append(bookIds, carts[i].BookID)
	}

	request := transactiondto.CreateTransactionRequest{
		BuyerID: userId,
		// BookID:  bookIds,
		Total:  total,
		Status: "pending",
	}

	transaction := models.Transaction{
		BuyerID: request.BuyerID,
		BookID:  request.BookID,
		Total:   request.Total,
		Status:  request.Status,
	}

	transaction, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, err = h.TransactionRepository.GetTransactionByID(transaction.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: transaction}
	json.NewEncoder(w).Encode(response)

}
