package models

type Programs struct {
	Title            string `json:"title"`
	Slug             string `json:"slug"`
	Description      string `json:"description"`
	Delete           bool   `json:"delete"`
	Create_at        string `json:"create_at"`
	Update_at        string `json:"update_at"`
	Delete_at        string `json:"delete_at"`
	Create_by        string `json:"create_by"`
	Update_by        string `json:"update_by"`
	Delete_by        string `json:"delete_by"`
	Image_destop     string `json:"image_destop"`
	Image_mobile     string `json:"image_mobile"`
	Level            string `json:"level"`
	Is_certification bool   `json:"is_certification"`
	Url_Logo         string `json:"url_logo"`
	Pic_name         string `json:"pic_name"`
	Pic_phone        string `json:"pic_phone"`
	Start_at         string `json:"start_at"`
	End_at           string `json:"end_at"`
	Is_active        bool   `json:"is_active"`
	Is_publish       bool   `json:"is_publish"`
}
