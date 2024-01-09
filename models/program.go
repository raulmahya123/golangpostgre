package models

import "gorm.io/gorm"

type Program struct {
	ID               uint   `gorm:"primarykey" json:"id"`
	Title            string `gorm:"type:varchar(255)" json:"title"`
	Slug             string `gorm:"type:varchar(255)" json:"slug"`
	Description      string `gorm:"type:text" json:"description"`
	Delete           bool   `gorm:"type:boolean" json:"delete"`
	Create_at        string `gorm:"null" json:"create_at"`
	Update_at        string `gorm:"null" json:"update_at"`
	Delete_at        string `gorm:"null" json:"delete_at"`
	Create_by        string `gorm:"type:varchar(255)" json:"create_by"`
	Update_by        string `gorm:"type:varchar(255)" json:"update_by"`
	Delete_by        string `gorm:"type:varchar(255)" json:"delete_by"`
	Image_destop     string `gorm:"type:varchar(255)" json:"image_destop"`
	Image_mobile     string `gorm:"type:varchar(255)" json:"image_mobile"`
	Level            int    `gorm:"type:int" json:"level"`
	Is_certification bool   `gorm:"type:boolean" json:"is_certification"`
	Url_Logo         string `gorm:"type:varchar(255)" json:"url_logo"`
	Pic_name         string `gorm:"type:varchar(255)" json:"pic_name"`
	Pic_phone        string `gorm:"type:varchar(255)" json:"pic_phone"`
	Start_at         string `gorm:"null"  json:"start_at"`
	End_at           string `gorm:"null"  json:"end_at"`
	Is_active        bool   `gorm:"type:boolean" json:"is_active"`
	Is_publish       bool   `gorm:"type:boolean" json:"is_publish"`
}

func MigrateProgram(db *gorm.DB) error {
	err := db.AutoMigrate(&Program{})
	return err
}
