package model

type User struct {
	UserID    int    `json:"user_id"`
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	MobileNo  string `json:"mobile_no"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
