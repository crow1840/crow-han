package service

import (
	userProto "crow-han/proto/user"
)

type ResetUserPasswordRequest struct {
	UserId      int64  `json:"user_id"`
	NewPassword string `json:"new_password"`
}

type GetUsersRequest struct {
	PageNum  int64  `json:"page_num"`
	PageSize int64  `json:"page_size"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

type GetUsersResponse struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
	Count    int64 `json:"count"`
	SysUsers []*userProto.GetUserResponse
}

type GetUserResponse struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type UpdateUsersInfoRequest struct {
	UserId int64  `json:"user_id"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}
