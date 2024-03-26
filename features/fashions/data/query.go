package data

import (
	"aifash-api/features/fashions"

	"gorm.io/gorm"
)

type FashionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) fashions.FashionDataInterface {
	return &FashionData{
		db: db,
	}
}
