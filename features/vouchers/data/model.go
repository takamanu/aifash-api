package data

import (
	"time"

	"gorm.io/gorm"
)

type VoucherStatus string

const (
	Used    VoucherStatus = "used"
	NotUsed VoucherStatus = "not_used"
	Expired VoucherStatus = "expired"
	Banned  VoucherStatus = "banned"
)

type Voucher struct {
	*gorm.Model
	VoucherName     string        `gorm:"unique;column:voucher_name" json:"voucher_name"`
	VoucherImageURL string        `gorm:"column:voucher_url_image" json:"voucher_url_image"`
	VoucherCode     string        `gorm:"unique;column:voucher_code" json:"voucher_code"`
	VoucherPrice    uint          `gorm:"column:voucher_price" json:"voucher_price"`
	Stock           uint          `gorm:"column:stock"`
	ExpiredIn       uint          `gorm:"column:expired_in"`
	UserVoucher     []UserVoucher `gorm:"foreignKey:VoucherID"`
}

type UserVoucher struct {
	*gorm.Model
	UserID         uint          `gorm:"column:user_id" json:"user_id"`
	VoucherID      uint          `gorm:"column:voucher_id"`
	Status         VoucherStatus `gorm:"column:voucher_status"`
	ExpirationDate *time.Time    `gorm:"column:expiration_date"`
}
