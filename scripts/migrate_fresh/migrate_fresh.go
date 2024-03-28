package main

import (
	"aifash-api/configs"
	"aifash-api/utils/database"

	"github.com/sirupsen/logrus"
)

func main() {

	var config = configs.InitConfig()

	db, err := database.InitDB(*config)
	if err != nil {
		logrus.Fatal("Cannot run database: ", err.Error())
	}

	database.MigrateWithDrop(db)
}
