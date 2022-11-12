package booksdto

type AddBookRequest struct {
	Title           string `json:"title" form:"title" gorm:"type:varchar(255)"`
	Author          string `json:"author" form:"author" gorm:"type:varchar(255)"`
	PublicationDate string `json:"publication_date" form:"publication_date" gorm:"type:varchar(255)"`
	Pages           int    `json:"pages" form:"publication_date" gorm:"type:int"`
	ISBN            string `json:"isbn" form:"isbn" gorm:"type:varchar(255)"`
	Price           int    `json:"price" form:"price" gorm:"type:varchar(255)"`
	About           string `json:"about" form:"about" gorm:"type:varchar(255)"`
	File            string `json:"file" form:"file" gorm:"type:varchar(255)"`
	Cover           string `json:"cover" form:"cover" gorm:"type:varchar(255)"`
}
type UpdateBookRequest struct {
	Title           string `json:"title" form:"title" gorm:"type:varchar(255)"`
	Author          string `json:"author" form:"author" gorm:"type:varchar(255)"`
	PublicationDate string `json:"publication_date" form:"publication_date" gorm:"type:varchar(255)"`
	Pages           int    `json:"pages" form:"publication_date" gorm:"type:int"`
	ISBN            string `json:"isbn" form:"isbn" gorm:"type:varchar(255)"`
	Price           int    `json:"price" form:"price" gorm:"type:varchar(255)"`
	About           string `json:"about" form:"about" gorm:"type:varchar(255)"`
	File            string `json:"file" form:"file" gorm:"type:varchar(255)"`
	Cover           string `json:"cover" form:"cover" gorm:"type:varchar(255)"`
}
