package data

import (
	"aifash-api/features/vouchers"
	"errors"

	"gorm.io/gorm"
)

type VoucherData struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) vouchers.VoucherDataInterface {
	return &VoucherData{
		db: db,
	}
}

func (vd *VoucherData) StoreVoucher(newData vouchers.Voucher) (*vouchers.Voucher, error) {
	var dbData = new(Voucher)
	dbData.VoucherName = newData.VoucherName
	dbData.VoucherImageURL = newData.VoucherImageURL
	dbData.VoucherCode = newData.VoucherCode
	dbData.VoucherPrice = newData.VoucherPrice
	dbData.Stock = newData.Stock
	dbData.ExpiredIn = newData.ExpiredIn

	if err := vd.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	dbDataReturned := vouchers.Voucher{
		VoucherName:     dbData.VoucherName,
		VoucherImageURL: dbData.VoucherImageURL,
		VoucherCode:     dbData.VoucherCode,
		VoucherPrice:    dbData.VoucherPrice,
		Stock:           dbData.Stock,
		ExpiredIn:       dbData.ExpiredIn,
	}

	return &dbDataReturned, nil
}
func (vd *VoucherData) GetAllVoucher() ([]vouchers.Voucher, error) {

	var vouchers []vouchers.Voucher

	if err := vd.db.Model(&Voucher{}).Where("deleted_at IS NULL").Scan(&vouchers).Error; err != nil {
		return nil, err
	}

	return vouchers, nil
}
func (vd *VoucherData) GetVoucherByID(id int) (*vouchers.Voucher, error) {
	var vouchers vouchers.Voucher

	if err := vd.db.Model(&Voucher{}).Where("id = ?", id).Where("deleted_at IS NULL").First(&vouchers).Error; err != nil {
		return nil, err
	}

	return &vouchers, nil
}
func (vd *VoucherData) GetVoucherByUserID(userID int) ([]vouchers.UserVoucher, error) {
	var userVoucher []vouchers.UserVoucher

	if err := vd.db.Model(&UserVoucher{}).Where("user_id = ?", userID).Where("deleted_at IS NULL").Scan(&userVoucher).Error; err != nil {
		return nil, err
	}

	return userVoucher, nil
}
func (vd *VoucherData) UpdateVoucherByID(id int, newData vouchers.Voucher) (bool, error) {
	_, err := vd.GetVoucherByID(id)

	if err != nil {
		return false, errors.New("voucher not found")
	}

	dbData := &Voucher{
		VoucherName:     newData.VoucherName,
		VoucherImageURL: newData.VoucherImageURL,
		VoucherCode:     newData.VoucherCode,
		VoucherPrice:    newData.VoucherPrice,
		Stock:           newData.Stock,
		ExpiredIn:       newData.ExpiredIn,
	}

	if err := vd.db.Model(&Voucher{}).Where("id = ?", id).Updates(dbData).Error; err != nil {
		return false, err
	}

	return true, nil
}
func (vd *VoucherData) DeleteVoucherByID(id int) (bool, error) {
	if err := vd.db.Delete(&Voucher{}, "id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}
