package data

import (
	"aifash-api/features/fashions"
	"errors"

	"gorm.io/gorm"
)

type FashionData struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) fashions.FashionDataInterface {
	return &FashionData{
		db: db,
	}
}

func (fd *FashionData) StoreFashion(newData fashions.Fashion) (*fashions.Fashion, error) {
	dbData := &Fashion{
		UserID:          newData.UserID,
		FashionName:     newData.FashionName,
		FashionPoints:   newData.FashionPoints,
		Status:          FashionStatus(OnProcess),
		FashionURLImage: newData.FashionURLImage,
	}

	if err := fd.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	dbDataReturned := fashions.Fashion{
		UserID:          dbData.UserID,
		FashionName:     dbData.FashionName,
		FashionPoints:   dbData.FashionPoints,
		Status:          string(FashionStatus(dbData.Status)),
		FashionURLImage: dbData.FashionURLImage,
	}

	return &dbDataReturned, nil
}
func (fd *FashionData) GetAllFashion() ([]fashions.Fashion, error) {
	var fashions []fashions.Fashion

	if err := fd.db.Model(&Fashion{}).Where("deleted_at IS NULL").Scan(&fashions).Error; err != nil {
		return nil, err
	}

	return fashions, nil
}
func (fd *FashionData) GetFashionByID(id int) (*fashions.Fashion, error) {
	var fashions fashions.Fashion

	if err := fd.db.Model(&Fashion{}).
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		Find(&fashions).Error; err != nil {
		return nil, err
	}

	return &fashions, nil
}
func (fd *FashionData) GetFashionByUserID(userID int) ([]fashions.Fashion, error) {
	var fashions []fashions.Fashion

	if err := fd.db.Model(&Fashion{}).
		Where("deleted_at IS NULL").
		Where("user_id = ?", userID).
		Scan(&fashions).Error; err != nil {
		return nil, err
	}

	return fashions, nil
}
func (fd *FashionData) UpdateFashionByID(id int, newData fashions.Fashion) (bool, error) {
	fashion, err := fd.GetFashionByID(id)

	if err != nil {
		return false, errors.New("fashion not found")
	}

	if fashion.Status == "accepted" {
		return false, errors.New("you cannot update an accepted fashion")
	}

	dbData := &Fashion{
		UserID:          newData.UserID,
		FashionName:     newData.FashionName,
		FashionPoints:   newData.FashionPoints,
		Status:          FashionStatus(newData.Status),
		FashionURLImage: newData.FashionURLImage,
	}

	if err := fd.db.Model(&Fashion{}).Where("id = ?", id).Updates(dbData).Error; err != nil {
		return false, err
	}

	return true, nil
}
func (fd *FashionData) DeleteFashionByID(id int) (bool, error) {
	if err := fd.db.Delete(&Fashion{}, "id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}
