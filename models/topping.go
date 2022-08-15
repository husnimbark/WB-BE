package models

import "time"

type Topping struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" form:"price" gorm:"type: int"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	ProductID int       `json:"product_id" form:"product_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ToppingResponse struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Price     int             `json:"price"`
	Image     string          `json:"image"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
}

type ToppingProductResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
	ProductID int    `json:"product_id"`
}

func (ToppingResponse) TableName() string {
	return "topping"
}

func (ToppingProductResponse) TableName() string {
	return "topping"
}
