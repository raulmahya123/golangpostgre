// models/user.go
package models

type User struct {
	ID        int    `json:"id"`
	Username  string ` json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	Create_at string `json:"create_at"`
	Create_by string `json:"create_by"`
	Update_by string `json:"update_by"`
	Delete_at bool   `json:"delete_at"`
	Delete_by string `json:"delete_by"`
}
