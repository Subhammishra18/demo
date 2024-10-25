package model

type LoginRequestPayload struct {
	MobileNumber string `json:"mobile_number"`
	Password     string `json:"password"`
}
