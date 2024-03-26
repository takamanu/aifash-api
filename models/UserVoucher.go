package models

import (
	"time"

	"gorm.io/gorm"
)

type UserVoucher struct {
	gorm.Model
	// VoucherID     uint `gorm:"primaryKey"`
	UserID         uint `gorm:"column:user_id" json:"user_id"`
	VoucherID      uint
	Voucher        Voucher    `gorm:"foreignKey:VoucherID"`
	Status         bool       `json:"status"`
	ExpirationDate *time.Time `json:"voucher_expired_date"`
}

func (UserVoucher) TableName() string {
	return "user_voucher"
}
