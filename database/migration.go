package database

import (
	"wayslink/models"
	"wayslink/pkg/postgre"
	"fmt"
)

func RunMigration() {
	err := postgre.DB.AutoMigrate(
		&models.User{},
		&models.Link{},
		&models.Sosmed{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}