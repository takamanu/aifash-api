package vouchers

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Voucher struct {
	VoucherName     string        `gorm:"column:voucher_name" json:"voucher_name"`
	VoucherImageURL string        `gorm:"column:voucher_url_image" json:"voucher_url_image"`
	VoucherCode     string        `gorm:"column:voucher_code" json:"voucher_code"`
	VoucherPrice    uint          `gorm:"column:voucher_price" json:"voucher_price"`
	Stock           uint          `gorm:"column:stock" json:"stock"`
	ExpiredIn       uint          `gorm:"column:expired_in" json:"expired_in"`
	UserVoucher     []UserVoucher `gorm:"foreignKey:VoucherID" json:"user_voucher"`
}

type UserVoucher struct {
	UserID         uint       `gorm:"column:user_id" json:"user_id"`
	VoucherID      uint       `gorm:"column:voucher_id" json:"voucher_id"`
	Status         string     `gorm:"column:voucher_status" json:"voucher_status"`
	ExpirationDate *time.Time `gorm:"column:expiration_date" json:"expiration_date"`
}

type VoucherHandlerInterface interface {
	StoreVoucher() echo.HandlerFunc
	GetAllVoucher() echo.HandlerFunc
	GetVoucherByID() echo.HandlerFunc
	GetVoucherByUserID() echo.HandlerFunc
	UpdateVoucherByID() echo.HandlerFunc
	DeleteVoucherByID() echo.HandlerFunc
}

type VoucherServiceInterface interface {
	StoreVoucher(newData Voucher) (*Voucher, error)
	GetAllVoucher() ([]Voucher, error)
	GetVoucherByID(id int) (*Voucher, error)
	GetVoucherByUserID(userID int) ([]UserVoucher, error)
	UpdateVoucherByID(id int, newData Voucher) (bool, error)
	DeleteVoucherByID(id int) (bool, error)
}

type VoucherDataInterface interface {
	StoreVoucher(newData Voucher) (*Voucher, error)
	GetAllVoucher() ([]Voucher, error)
	GetVoucherByID(id int) (*Voucher, error)
	GetVoucherByUserID(userID int) ([]UserVoucher, error)
	UpdateVoucherByID(id int, newData Voucher) (bool, error)
	DeleteVoucherByID(id int) (bool, error)
}
