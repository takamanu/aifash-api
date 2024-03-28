package handler

type InputRequestVoucher struct {
	VoucherName     string `json:"voucher_name" form:"voucher_name"`
	VoucherImageURL string `json:"voucher_url_image" form:"voucher_url_image"`
	VoucherCode     string `json:"voucher_code" form:"voucher_code"`
	VoucherPrice    uint   `json:"voucher_price" form:"voucher_price"`
	Stock           uint   `json:"stock" form:"stock"`
	ExpiredIn       uint   `json:"expired_in" form:"expired_in"`
}

type InputRequestUserVoucher struct {
	VoucherID uint `json:"voucher_id" form:"voucher_id"`
}

type UpdateRequestUserVoucher struct {
	Status string `json:"status" form:"status"`
}
