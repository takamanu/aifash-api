package fashions

import "github.com/labstack/echo/v4"

type Fashion struct {
	ID              uint   `json:"id,omitempty"`
	UserID          uint   `gorm:"column:user_id" json:"user_id"`
	FashionName     string `gorm:"column:fashion_name" json:"fashion_name"`
	FashionPoints   int    `gorm:"column:fashion_points" json:"fashion_points"`
	Status          string `gorm:"column:status" json:"status"`
	FashionURLImage string `gorm:"column:fashion_url_image" json:"fashion_url_image"`
}

type FashionDetailed struct {
	UserID          uint   `gorm:"column:user_id" json:"user_id"`
	UserName        string `json:"name" gorm:"-"`
	FashionName     string `gorm:"column:fashion_name" json:"fashion_name"`
	FashionPoints   int    `gorm:"column:fashion_points" json:"fashion_points"`
	Status          string `gorm:"column:status" json:"status"`
	FashionURLImage string `gorm:"column:fashion_url_image" json:"fashion_url_image"`
}

type FashionHandlerInterface interface {
	StoreFashion() echo.HandlerFunc
	GetAllFashion() echo.HandlerFunc
	GetFashionByID() echo.HandlerFunc
	GetFashionByUserID() echo.HandlerFunc
	UpdateFashionByID() echo.HandlerFunc
	DeleteFashionByID() echo.HandlerFunc
}

type FashionServiceInterface interface {
	StoreFashion(newData Fashion) (*Fashion, error)
	GetAllFashion() ([]Fashion, error)
	GetFashionByID(id int) (*Fashion, error)
	GetFashionByUserID(userID int) ([]Fashion, error)
	UpdateFashionByID(id int, newData Fashion) (bool, error)
	DeleteFashionByID(id int) (bool, error)
}

type FashionDataInterface interface {
	StoreFashion(newData Fashion) (*Fashion, error)
	GetAllFashion() ([]Fashion, error)
	GetFashionByID(id int) (*Fashion, error)
	GetFashionByUserID(userID int) ([]Fashion, error)
	UpdateFashionByID(id int, newData Fashion) (bool, error)
	DeleteFashionByID(id int) (bool, error)
}
