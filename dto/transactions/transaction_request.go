package transactiondto

type CreateTransactionRequest struct {
	BuyerID int    `json:"buyer_id"`
	BookID  []int  `json:"book_id"`
	Total   int    `json:"total"`
	Status  string `json:"status"`
}
