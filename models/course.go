package models

type Course struct {
	Category_id int    `json:"category_id"`
	Owner_id    int    `json:"owner_id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Privacy     string `json:"privacy"`
	Start_at    string `json:"start_at"`
	End_at      string `json:"end_at"`
	Image       string `json:"image"`
	Delete      bool   `json:"delete"`
	Create_at   string `json:"create_at"`
	Update_at   string `json:"update_at"`
	Delete_at   string `json:"delete_at"`
	Create_by   string `json:"create_by"`
	Update_by   string `json:"update_by"`
	Delete_by   string `json:"delete_by"`
	Certificate bool   `json:"certificate"`
	Level       string `json:"level"`
	Price       int    `json:"price"`
	Status_enum string `json:"status_enum"`
}

type CourseRequest struct {
	Category_id int    `json:"category_id"`
	Owner_id    int    `json:"owner_id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Privacy     string `json:"privacy"`
	Start_at    string `json:"start_at"`
	End_at      string `json:"end_at"`
	Image       string `json:"image"`
	Certificate bool   `json:"certificate"`
	Level       string `json:"level"`
	Price       int    `json:"price"`
	Status_enum string `json:"status_enum"`
}
