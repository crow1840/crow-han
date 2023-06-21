package service

type UpdateUserSelfInfoRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UpdateUserPasswordRequest struct {
	//UserId      int64  `json:"user_id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
