package api

type Order struct {
	Name          string
	Content       string
	Purchaser     *User
	CreatedAtDate string
	Status        string
}
