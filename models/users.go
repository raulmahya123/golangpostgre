package models

import (
	"gorm.io/gorm"
)

type Users struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Username *string `gorm:"not null" json:"username"`
	Password *string `gorm:"not null" json:"-"`
	Role     *string `gorm:"not null" json:"role"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
