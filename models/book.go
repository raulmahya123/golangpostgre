package models

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Delete_at bool   `json:"delete_at"`
	Create_at string `json:"create_at"`
	Create_by string `json:"create_by"`
	Delete_by string `json:"delete_by"`
	Update_at string `json:"create_at"`
	Update_by bool   `json:"update_by"`
}
type BookRequest struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}
