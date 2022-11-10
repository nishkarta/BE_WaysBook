package models

import "time"

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	FullName string `json:"fullName" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type:varchar(255)"`
	Address  string `json:"address"`
	Role     string `json:"role" gorm:"type:varchar(255)"`
	Image    string `json:"image" gorm:"type:varchar(255)"`
	// Transaction []TransactionResponse `json:"transaction"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}
