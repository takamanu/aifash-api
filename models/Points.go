package models

import (
	"gorm.io/gorm"
	"time"
)

type Point struct {
	gorm.Model
    PointID       uint `gorm:"primaryKey"`
    UserID        uint
    User          User `gorm:"foreignKey:UserID"`
    EarnedPoints  int
    Timestamp     time.Time
}