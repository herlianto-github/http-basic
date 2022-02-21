package users

import "httpBasic/entities"

type UsersInterface interface {
	Gets() ([]entities.User, error)
}
