package routes

import (
	"aifash-api/configs"
	"aifash-api/features/users"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Group, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/register", uh.RegisterCustomer())
	e.POST("/admin/register", uh.Register())
	// e.POST("/login", uh.LoginCustomer())
	e.POST("/login", uh.Login())
	e.POST("/forget-password", uh.ForgetPasswordWeb())
	e.POST("/forget-password/verify", uh.ForgetPasswordVerify())
	e.POST("/reset-password", uh.ResetPassword())
	// e.POST("/refresh-token", uh.RefreshToken(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/admin/update", uh.UpdateProfile(), echojwt.JWT([]byte(cfg.Secret)))
	// e.GET("/user/profile", uh.GetProfile(), echojwt.JWT([]byte(cfg.Secret)))
}

// // blogs routes
// e.POST("/blogs", controllers.CreateBlogController)
// e.GET("/blogs/:id", controllers.GetBlogController)
// e.PUT("/blogs/:id", controllers.UpdateBlogController)
// e.DELETE("/blogs/:id", controllers.DeleteBlogController)

// e.POST("/fashion", CreateFashionItemHandler)
// e.GET("/fashion", controllers.GetFashionItemsController)
// e.GET("/fashion/:id", controllers.GetFashionItemByIDController)
// e.PATCH("/fashion/:id", controllers.ChooseResponseFashion)
// e.PUT("/fashion/:id", controllers.UpdateFashionItemController)

// e.POST("/voucher", CreateVoucherItemHandler)
// e.GET("/voucher", controllers.GetVouchers)
// e.GET("/voucher/:id", controllers.GetVoucherByID)
// e.PATCH("/voucher/:id", controllers.UpdateVoucherController)

// e.POST("/voucher/apply", controllers.CreateUserVoucher)
// e.PATCH("/voucher/use/:id", controllers.MarkVoucherAsUsed)
