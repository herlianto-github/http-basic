package users

import (
	"database/sql"
	"errors"
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
	config := configs.GetConfig()
	dbs := utils.InitDB(config)
	defer dbs.Close()

	_, err := dbs.Exec("insert into users values (?, ?, ?, ?)", &newUser.ID, &newUser.Name, &newUser.Email, &newUser.Password)
	if err != nil {
		return newUser, errors.New("ERROR REGISTER USER")
	}

	return newUser, nil
}

func (ur *UserRepository) Gets() ([]entities.User, error) {
	users := []entities.User{}

	config := configs.GetConfig()
	dbs := utils.InitDB(config)
	defer dbs.Close()

	res, err := dbs.Query("select * from users")
	if err != nil {
		return users, errors.New("ERROR QUERY")
	}
	defer res.Close()

	for res.Next() {
		var each = entities.User{}
		var err = res.Scan(&each.ID, &each.Name, &each.Email, &each.Password)
		if err != nil {
			return users, errors.New("ERROR SCAN")
		}
		users = append(users, each)
	}

	if err = res.Err(); err != nil {
		return users, errors.New("ERROR GETS USER")
	}

	return users, nil
}
