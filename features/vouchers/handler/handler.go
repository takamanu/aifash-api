package handler

import (
	"aifash-api/features/users"
	"aifash-api/features/vouchers"
	"aifash-api/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type VoucherHandler struct {
	vs  vouchers.VoucherServiceInterface
	us  users.UserServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(vs vouchers.VoucherServiceInterface, us users.UserServiceInterface, jwt helper.JWTInterface) vouchers.VoucherHandlerInterface {
	return &VoucherHandler{
		vs:  vs,
		us:  us,
		jwt: jwt,
	}
}

func (vh *VoucherHandler) StoreVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = new(InputRequestVoucher)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Bind input Error", nil))
		}

		code, _ := helper.Generate(helper.RandomString(15))
		voucherCode := strings.ToUpper("VCR" + code)

		res, err := vh.vs.StoreVoucher(
			vouchers.Voucher{
				VoucherName:     input.VoucherName,
				VoucherImageURL: input.VoucherImageURL,
				VoucherCode:     voucherCode,
				VoucherPrice:    input.VoucherPrice,
				Stock:           input.Stock,
				ExpiredIn:       input.ExpiredIn,
			})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "success create voucher", res))
	}
}
func (vh *VoucherHandler) ClaimVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := vh.jwt.GetID(c)

		var input = new(InputRequestUserVoucher)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Bind input Error", nil))
		}

		voucher, err := vh.vs.GetVoucherByID(int(input.VoucherID))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		user, err := vh.us.GetProfile(int(userID))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		if user.Points < voucher.VoucherPrice {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "insufficient points", nil))
		}

		res, err := vh.vs.ClaimVoucher(
			vouchers.UserVoucher{
				UserID:    userID,
				VoucherID: input.VoucherID,
			})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		_, err = vh.us.DeductPoints(int(userID), int(voucher.VoucherPrice))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success claim voucher", res))
	}
}

func (vh *VoucherHandler) UpdateClaimedVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var input = new(UpdateRequestUserVoucher)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Bind input Error", nil))
		}

		_, err := vh.vs.UpdateClaimedVoucher(id, vouchers.UserVoucher{
			Status: input.Status,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success update claimed voucher", nil))
	}
}
func (vh *VoucherHandler) GetAllVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := vh.vs.GetAllVoucher()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get vouchers", res))
	}

}
func (vh *VoucherHandler) GetVoucherByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		res, err := vh.vs.GetVoucherByID(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get voucher", res))
	}
}
func (vh *VoucherHandler) GetUserVoucherByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		res, err := vh.vs.GetUserVoucherByID(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get user voucher", res))
	}
}
func (vh *VoucherHandler) GetVoucherByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := vh.jwt.GetID(c)

		res, err := vh.vs.GetVoucherByID(int(userID))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get voucher user", res))
	}
}
func (vh *VoucherHandler) UpdateVoucherByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		var input = new(InputRequestVoucher)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Bind input Error", nil))
		}

		_, err := vh.vs.UpdateVoucherByID(id,
			vouchers.Voucher{
				VoucherName:     input.VoucherName,
				VoucherImageURL: input.VoucherImageURL,
				VoucherPrice:    input.VoucherPrice,
				Stock:           input.Stock,
				ExpiredIn:       input.ExpiredIn,
			})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success update voucher", nil))
	}
}
func (vh *VoucherHandler) DeleteVoucherByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := vh.vs.DeleteVoucherByID(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse(true, "success delete voucher", nil))
	}
}
