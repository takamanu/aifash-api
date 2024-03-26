package data

import "gorm.io/gorm"

type FashionStatus string

const (
	OnProcess FashionStatus = "on_process"
	Accepted  FashionStatus = "accepted"
	Denied    FashionStatus = "denied"
)

type Fashion struct {
	*gorm.Model
	UserID          uint          `gorm:"column:user_id" json:"user_id"`
	FashionName     string        `gorm:"column:fashion_name" json:"fashion_name"`
	FashionPoints   int           `gorm:"column:fashion_points" json:"fashion_points"`
	Status          FashionStatus `gorm:"column:status" json:"status"`
	FashionURLImage string        `gorm:"column:fashion_url_image" json:"fashion_url_image"`
}
