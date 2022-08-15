package authdto

type LoginResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
