package routes

import (
	"aifash-api/configs"
	"aifash-api/features/fashions"
	"aifash-api/features/users"
	"aifash-api/features/vouchers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Group, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/register", uh.Register())
	// e.POST("/admin/register", uh.Register())
	// e.POST("/login", uh.LoginCustomer())
	e.POST("/login", uh.Login())
	e.POST("/forget-password", uh.ForgetPasswordWeb())
	e.POST("/forget-password/verify", uh.ForgetPasswordVerify())
	e.POST("/reset-password", uh.ResetPassword())
	// e.POST("/refresh-token", uh.RefreshToken(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/admin/update", uh.UpdateProfile(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/user/profile", uh.GetProfile(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteFashion(e *echo.Group, fh fashions.FashionHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/fashion", fh.StoreFashion())
	e.POST("/upload", fh.UploadFile())
	e.GET("/fashion", fh.GetAllFashion())
	e.GET("/fashion/:id", fh.GetFashionByID())
	e.GET("/fashion/user", fh.GetFashionByUserID())
	e.PUT("/fashion/:id", fh.UpdateFashionByID())
	e.DELETE("/fashion/:id", fh.DeleteFashionByID())
}

func RouteVoucher(e *echo.Group, vh vouchers.VoucherHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/voucher", vh.StoreVoucher())
	e.GET("/voucher", vh.GetAllVoucher())
	e.GET("/voucher/:id", vh.GetVoucherByID())
	e.GET("/voucher/user", vh.GetVoucherByUserID())
	e.PUT("/voucher/:id", vh.UpdateVoucherByID())
	e.DELETE("/voucher/:id", vh.DeleteVoucherByID())

	e.POST("/user-voucher", vh.ClaimVoucher())
	e.GET("/user-voucher/:id", vh.GetUserVoucherByID())
	e.PUT("/user-voucher/:id", vh.UpdateClaimedVoucher())
}
