package users

import "httpBasic/entities"

type UsersInterface interface {
	Gets() ([]entities.User, error)
	AddUser(newUser entities.User) (entities.User, error)
}
