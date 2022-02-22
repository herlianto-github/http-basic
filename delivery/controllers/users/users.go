package users

import (
	"encoding/json"
	"httpBasic/delivery/common"
	"httpBasic/entities"
	"httpBasic/repository/users"
	"net/http"
)

type UsersController struct {
	Repo users.UsersInterface
}

func NewUsersControllers(usrep users.UsersInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

func (uscon UsersController) Gets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":

			if res, err := uscon.Repo.Gets(); err != nil {
				http.Error(w, common.NewNotFoundResponse().Message, common.NewBadRequestResponse().Code)
			} else {
				common.SetHeaderResponse(w, r)
				json.NewEncoder(w).Encode(common.NewSuccessOperationResponse(res))
			}

		case "POST":

			param := json.NewDecoder(r.Body)
			user := entities.User{}
			if err := param.Decode(&user); err != nil {
				http.Error(w, common.NewBadRequestResponse().Message, common.NewBadRequestResponse().Code)
			} else {

				newUser := entities.User{
					Name: user.Name, Email: user.Email, Password: user.Password,
				}

				if res, err := uscon.Repo.Register(newUser); err != nil {
					http.Error(w, common.NewInternalServerErrorResponse().Message, common.NewInternalServerErrorResponse().Code)
				} else {
					common.SetHeaderResponse(w, r)
					json.NewEncoder(w).Encode(common.NewSuccessOperationResponse(res))
				}
			}

		default:
			http.Error(w, common.NewBadRequestResponse().Message, common.NewBadRequestResponse().Code)
		}

	}
}
