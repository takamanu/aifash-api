package database

import (
	"fmt"

	DataUser "aifash-api/features/users/data"

	"gorm.io/gorm"
)

func MigrateWithDrop(db *gorm.DB) {

	// USER DATA MANAGEMENT \\
	db.AutoMigrate(DataUser.User{})
	db.AutoMigrate(DataUser.UserResetPass{})
	fmt.Println("[MIGRATION] Success creating user and regions")

	// Fashion DATA MANAGEMENT \\
	// 	DB.AutoMigrate(&models.Fashion{})

	// Voucher DATA MANAGEMENT \\
	// 	DB.AutoMigrate(&models.Point{})
	// 	DB.AutoMigrate(&models.Voucher{})
	// 	DB.AutoMigrate(&models.UserVoucher{})

}
