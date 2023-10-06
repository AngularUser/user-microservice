package model

type User struct {
	UserID string `json:"UserID"`
	Name   string `json:"Name"`
	Email  string `json:"Email"`
	DOB    string `json:"DOB"`
	// Add other fields as needed
}
