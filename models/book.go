package models

import "time"

type Book struct {
	ID              int       `json:"id" gorm:"primary_key:auto_increment"`
	Title           string    `json:"title" gorm:"type:varchar(255)"`
	Author          string    `json:"author" gorm:"type:varchar(255)"`
	PublicationDate string    `json:"publication_date" gorm:"type:varchar(255)"`
	Pages           int       `json:"pages" gorm:"type:int"`
	ISBN            string    `json:"isbn" gorm:"type:varchar(255)"`
	Price           int       `json:"price" gorm:"type:int"`
	About           string    `json:"about" gorm:"type:varchar(255)"`
	File            string    `json:"file" gorm:"type:varchar(255)"`
	Cover           string    `json:"cover" gorm:"type:varchar(255)"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

type UserBooksResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication_date"`
	Pages           int    `json:"pages"`
	ISBN            string `json:"isbn"`
	Price           int    `json:"price"`
	About           string `json:"about"`
	File            string `json:"file"`
	Cover           string `json:"cover"`
}

func (UserBooksResponse) TableName() string {
	return "books"
}
