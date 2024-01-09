package models

import "gorm.io/gorm"

type Courses struct {
	ID          uint   `gorm:"primary key;autoIncrement" json:"id"`
	Category_id int    `gorm:"not null" json:"category_id"`
	Owner_id    int    `gorm:"not null" json:"owner_id"`
	Title       string `gorm:"not null" json:"title"`
	Slug        string `gorm:"not null" json:"slug"`
	Description string `gorm:"not null" json:"description"`
	Privacy     string `gorm:"not null" json:"privacy"`
	Start_at    string `gorm:"not null" json:"start_at"`
	End_at      string `gorm:"not null" json:"end_at"`
	Image       string `gorm:"not null" json:"image"`
	Delete      bool   `gorm:"default:false" json:"delete"`
	Create_at   string `gorm:"null" json:"create_at"`
	Update_at   string `gorm:"null" json:"update_at"`
	Delete_at   string `gorm:"null" json:"delete_at"`
	Create_by   string `gorm:"null" json:"create_by"`
	Update_by   string `gorm:"null" json:"update_by"`
	Delete_by   string `gorm:"null" json:"delete_by"`
	Certificate bool   `gorm:"default:false" json:"certificate"`
	Level       string `gorm:"not null" json:"level"`
	Price       int    `gorm:"not null" json:"price"`
	Status_enum string `gorm:"not null" json:"status_enum"`
}

func MigrateCourse(db *gorm.DB) error {
	err := db.AutoMigrate(&Courses{})
	return err
}
