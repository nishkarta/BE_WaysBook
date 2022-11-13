package models

import "time"

type Transaction struct {
	ID      int                 `json:"id" gorm:"primary_key:auto_increment"`
	BuyerID int                 `json:"buyer_id"`
	Buyer   UserProfileResponse `json:"buyer" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Status  string              `json:"status"`
	Total   int                 `json:"total"`
	// Books []UserBooksResponse `json:"books"`
	BookID    []int               `json:"book_id" form:"book_id" gorm:"type:int"`
	Book      []UserBooksResponse `json:"book"  gorm:"many2many:book_cart"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}
