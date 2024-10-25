package model

type SignupRequestPayload struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Sex         string `json:"sex"`
	Password    string `json:"password"`
	MobileNumber string `json:"mobile_number"` // Added for signup
}
