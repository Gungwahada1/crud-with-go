package entities

import "time"


type User struct {
	Id int64 `json:"id"`
	No int `json:"no"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	RoleId string `json:"role_id"`
	ReligionId string `json:"religion_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}