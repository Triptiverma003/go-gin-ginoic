package model

import "time"

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"column:email;unique"`
	Password  string         `json:"password" gorm:"column:password;"`
	CreatedAt time.Time 	 `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
}