package models

import "time"

type Cart struct {
	ID        int                 `json:"id" gorm:"primaryKey;autoIncrement"`
	Price     int                 `json:"price"`
	BookID    int                 `json:"bookId"`
	Book      UserBooksResponse   `json:"book" gorm:"foreignKey:BookID;references:ID"`
	UserID    int                 `json:"userId"`
	User      UserProfileResponse `json:"user"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}
