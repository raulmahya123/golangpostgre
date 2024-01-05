package models

import "gorm.io/gorm"

type Books struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `gorm:"not null" json:"author"`
	Title     *string `gorm:"not null" json:"title"`
	Publisher *string `gorm:"not null" json:"publisher"`
	Delete_at bool    `gorm:"default:false" json:"delete_at"`
	Create_at string  `gorm:"not" json:"create_at"`
	Create_by string  `gorm:"not null" json:"create_by"`
	Delete_by string  `gorm:"-" json:"delete_by,omitempty"`
	Update_at string  `gorm:"not" json:"update_at"`
	Update_by bool    `gorm:"default:false" json:"update_by"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
