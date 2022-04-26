package entities

type Guests []Guest

type Guest struct {
	ID            int    `json:"id" param:"id"`
	FirstName     string `json:"first_name"`
	FirstNameKana string `json:"first_name_kana"`
	LastName      string `json:"last_name"`
	LastNameKana  string `json:"last_name_kana"`
	Gender        int    `json:"gender"`
	Birthday      string `json:"birthday"`
	Tel           int    `json:"tel"`
	Email         string `json:"email"`
}
