package service

import "aifash-api/features/vouchers"

type VoucherService struct {
	vd vouchers.VoucherDataInterface
}

func NewService(vd vouchers.VoucherDataInterface) vouchers.VoucherServiceInterface {
	return &VoucherService{
		vd: vd,
	}
}

func (vs *VoucherService) StoreVoucher(newData vouchers.Voucher) (*vouchers.Voucher, error) {
	res, err := vs.vd.StoreVoucher(newData)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VoucherService) ClaimVoucher(newData vouchers.UserVoucher) (*vouchers.UserVoucher, error) {
	res, err := vs.vd.ClaimVoucher(newData)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VoucherService) UpdateClaimedVoucher(id int, newData vouchers.UserVoucher) (bool, error) {
	_, err := vs.vd.UpdateClaimedVoucher(id, newData)

	if err != nil {
		return false, err
	}

	return true, nil

}
func (vs *VoucherService) GetAllVoucher() ([]vouchers.Voucher, error) {
	res, err := vs.vd.GetAllVoucher()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VoucherService) GetVoucherByID(id int) (*vouchers.Voucher, error) {
	res, err := vs.vd.GetVoucherByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VoucherService) GetUserVoucherByID(id int) (*vouchers.UserVoucher, error) {
	res, err := vs.vd.GetUserVoucherByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VoucherService) GetVoucherByUserID(userID int) ([]vouchers.UserVoucher, error) {
	res, err := vs.vd.GetVoucherByUserID(userID)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (vs *VoucherService) UpdateVoucherByID(id int, newData vouchers.Voucher) (bool, error) {

	_, err := vs.vd.UpdateVoucherByID(id, newData)

	if err != nil {
		return false, err
	}

	return true, nil
}
func (vs *VoucherService) DeleteVoucherByID(id int) (bool, error) {
	_, err := vs.vd.DeleteVoucherByID(id)

	if err != nil {
		return false, err
	}

	return true, nil
}
