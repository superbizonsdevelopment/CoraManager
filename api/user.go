package api

type User struct {
	Name        string
	Password    string
	Email       string
	Permissions string
	Orders      []*Order
}
