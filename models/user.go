package models

import "time"

type User struct {
	ID        int             `json:"id"`
	FullName  string          `json:"fullname" gorm:"type: varchar(255)"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"-" gorm:"type: varchar(255)"`
	Image     string          `json:"image" gorm:"type: varchar(255)"`
	Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
