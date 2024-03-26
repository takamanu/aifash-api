package models

import (
	"gorm.io/gorm"
)

type FashionStatus string

const (
	OnProcess FashionStatus = "on_process"
	Accepted  FashionStatus = "accepted"
	Denied    FashionStatus = "denied"
)

type Fashion struct {
	gorm.Model
	UserID     uint          `gorm:"column:user_id" json:"user_id"`
	ItemType   string        `json:"fashion_name"`
	ItemPoints int           `json:"fashion_points"`
	ItemStatus FashionStatus `json:"status"`
	ImageUrl   string        `json:"fashion_url_image"`
}
