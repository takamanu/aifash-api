package controllers

import (
	"curdusers/configs"
	"curdusers/helper"
	m "curdusers/middlewares"
	"curdusers/models"
	"fmt"

	loger "log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersController(c echo.Context) error {
	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		log.Errorf("Failed to get users: %s", err.Error())
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "failed to get user", nil))

	}
	return c.JSON(http.StatusOK, helper.FormatResponse(true, "success", users))
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Errorf("Invalid id: %s", c.Param("id"))
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "invalid id", nil))
	}

	var user models.User
	if err := configs.DB.First(&user, id).Error; err != nil {
		log.Errorf("Failed to get user with id %d: %s", id, err.Error())
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
	}

	var fashion []models.Fashion
	if err := configs.DB.Where("user_id = ?", user.ID).Find(&fashion).Error; err != nil {
		log.Errorf("Failed to get fashion for user with id %d: %s", user.ID, err.Error())
		return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
	}

	var userVouchers []models.UserVoucher
	if err := configs.DB.Preload("Voucher").Where("user_id = ?", user.ID).Find(&userVouchers).Error; err != nil {
		log.Errorf("Failed to get user_vouchers for user with id %d: %s", user.ID, err.Error())
		return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
	}

	if len(fashion) == 0 {
		fashion = []models.Fashion{}
	}
	if len(userVouchers) == 0 {
		userVouchers = []models.UserVoucher{}
	}

	user.Fashion = fashion
	user.UserVoucher = userVouchers

	return c.JSON(http.StatusOK, helper.FormatResponse(true, "sucess", user))
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		log.Errorf("Failed to bind request: %s", err.Error())
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
	}

	user.Password = string(hashedPassword)
	user.Points = 0
	loger.Println(user)

	if err := configs.DB.Create(&user).Error; err != nil {
		log.Errorf("Failed to create user: %s", err.Error())
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.FormatResponse(true, "success create", user))
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Errorf("Invalid id: %s", c.Param("id"))
		return c.JSON(http.StatusBadRequest, "Invalid id")
	}

	var user models.User
	if err := configs.DB.First(&user, id).Error; err != nil {
		log.Errorf("Failed to get user with id %d: %v", id, err)
		return c.JSON(http.StatusNotFound, "User not found")
	}

	if err := configs.DB.Delete(&user).Error; err != nil {
		log.Errorf("Failed to delete user with id %d: %v", id, err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.JSON(http.StatusNoContent, helper.FormatResponse(true, "success delete", nil))
}

func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	var user models.User
	if err := configs.DB.First(&user, userId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "user not found")
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newPassword := c.FormValue("password")
	if newPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to encrypt password")
		}
		user.Password = string(hashedPassword)
	}

	if err := configs.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.FormatResponse(true, "success updated", nil))
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := configs.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to login",
			"error":   err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.FormValue("password"))); err != nil {
		fmt.Println("pass :", c.FormValue("password"))
		fmt.Println("err :", err)
		return c.JSON(http.StatusUnauthorized, "invalid email or password")
	}

	fmt.Println("pass :", c.FormValue("password"))

	token, err := m.CreateToken(int(user.ID), user.Name, int(user.Role))
	fmt.Printf("UserID: %v, UserName: %v, UserRole: %v", user.ID, user.Name, user.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Role:  user.Role,
		Email: user.Email,
		Token: token}

	return c.JSON(http.StatusOK, helper.FormatResponse(true, "success login", userResponse))
}
