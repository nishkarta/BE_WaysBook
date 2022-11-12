package authdto

type LoginResponse struct {
	ID       int    `gorm:"type:int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type:varchar(255)" json:"role"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Image    string `gorm:"type:varchar(255)" json:"image"`
}

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type: varchar(50)"  json:"role"`
}
