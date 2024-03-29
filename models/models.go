package models

type User struct {
	ID        int32  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	createdAt string `json:"created_at"`
}
