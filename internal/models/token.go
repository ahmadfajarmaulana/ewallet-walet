package models

type TokenData struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}
