package users

import (
	"database/sql"
	"errors"
	"fmt"
	"httpBasic/configs"
	"httpBasic/entities"
	"httpBasic/utils"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Register(newUser entities.User) (entities.User, error) {
	return newUser, nil
}

func (ur *UserRepository) Gets() ([]entities.User, error) {
	users := []entities.User{}

	config := configs.GetConfig()
	dbs := utils.InitDB(config)
	defer dbs.Close()

	res, err := dbs.Query("select id,name,email,password from users")
	if err != nil {
		fmt.Println("Error Gets Users", err.Error())
	}
	defer res.Close()

	for res.Next() {
		var each = entities.User{}
		var err = res.Scan(&each.ID, &each.Name, &each.Email, &each.Password)
		if err != nil {
			fmt.Println("Error Scan", err.Error())
			return users, errors.New("ERROR GETS USER")
		}
		users = append(users, each)

	}

	if err = res.Err(); err != nil {
		fmt.Println("Error Res", err.Error())
		return users, errors.New("ERROR GETS USER")
	}

	return users, nil
}
