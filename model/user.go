package model

type User struct {
	UserID string `json:"UserID"`
	Name   string `json:"Name" validate:"required,username"`
	Email  string `json:"Email" validate:"required,email"`
	DOB    string `json:"DOB"`
}
