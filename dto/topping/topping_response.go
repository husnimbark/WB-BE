package toppingdto

type ToppingResponse struct {
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Desc  string `json:"desc" gorm:"type:text" form:"desc"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
}
