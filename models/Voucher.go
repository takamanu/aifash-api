package models

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	// VoucherID     uint `gorm:"primaryKey"`
	VoucherImageUrl string    `json:"voucher_url_image"`
	VoucherCode     string    `gorm:"unique" json:"voucher_code"`
	VoucherName     string    `gorm:"unique" json:"voucher_name"`
	VoucherValue    uint      `json:"voucher_price"`
	ExpirationDate  time.Time `json:"voucher_expired_date"`
}

func (Voucher) TableName() string {
	return "voucher"
}
