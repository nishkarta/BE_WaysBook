package cartdto

type CreateCartRequest struct {
	BookID int `json:"bookId"`
	UserID int `json:"userId"`
	Price  int `json:"price"`
}
