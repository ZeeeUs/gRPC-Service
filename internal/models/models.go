package models

type Account struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint32 `json:"age"`
}
