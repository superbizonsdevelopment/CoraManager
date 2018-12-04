package api

type Order struct {
	ID            int
	Name          string
	Content       string
	Purchaser     *User
	CreatedAtDate string
	Status        string
	Deadline      string
}
