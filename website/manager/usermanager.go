package manager

import (
	"strings"

	"github.com/superbizonsdevelopment/CoraManager/api"
)

func CreateUser(usr *api.User) error {
	var sb strings.Builder
	for _, order := range usr.Orders {
		sb.WriteString(string(order.ID))
	}
	_, err := Base.Exec("INSERT INTO users (name, password, email, permissions, orders) VALUES ($1, $2, $3, $4, $5)", usr.Name, usr.Password, usr.Email, usr.Permissions, sb.String())
	if err != nil {
		return err
	}
	return nil
}

/*
func DeleteUser(user *api.User) error {
  return nil
}

func GetAllUsers() ([]*api.User, error) {

}

func GetUserByID() {}
*/
