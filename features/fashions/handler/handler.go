package handler

import (
	"aifash-api/features/fashions"
	"aifash-api/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FashionHandler struct {
	fs  fashions.FashionServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(fs fashions.FashionServiceInterface, jwt helper.JWTInterface) fashions.FashionHandlerInterface {
	return &FashionHandler{
		fs:  fs,
		jwt: jwt,
	}
}

func (fh *FashionHandler) StoreFashion() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := fh.jwt.GetID(c)

		var input = new(InputRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Bind input Error", nil))
		}

		res, err := fh.fs.StoreFashion(
			fashions.Fashion{
				UserID:          userID,
				FashionName:     input.FashionName,
				FashionPoints:   input.FashionPoints,
				FashionURLImage: input.FashionURLImage,
			})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "success create data", res))
	}
}
func (fh *FashionHandler) GetAllFashion() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := fh.fs.GetAllFashion()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get data", res))

	}
}
func (fh *FashionHandler) GetFashionByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		res, err := fh.fs.GetFashionByID(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get data", res))

	}
}
func (fh *FashionHandler) GetFashionByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID, _ := fh.jwt.GetID(c)

		res, err := fh.fs.GetFashionByUserID(int(userID))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get data", res))

	}
}
func (fh *FashionHandler) UpdateFashionByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var input = new(InputRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Bind input Error", nil))
		}

		_, err := fh.fs.UpdateFashionByID(id,
			fashions.Fashion{
				FashionName:     input.FashionName,
				FashionPoints:   input.FashionPoints,
				Status:          input.Status,
				FashionURLImage: input.FashionURLImage,
			})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success update data", nil))
	}
}
func (fh *FashionHandler) DeleteFashionByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := fh.fs.DeleteFashionByID(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success delete data", nil))
	}
}
