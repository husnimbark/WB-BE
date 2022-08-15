package productdto

type ProductResponse struct {
	Title    string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Price    int    `json:"price" form:"price" gorm:"type: int"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID   int    `json:"user_id" gorm:"type: int"`
	TopingID int    `json:"toping_id gorm"`
}
