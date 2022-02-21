package users

import (
	"encoding/json"
	"httpBasic/delivery/common"
	"httpBasic/entities"
	"httpBasic/repository/users"
	"net/http"
	// "google.golang.org/genproto/googleapis/cloud/common"
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
				http.Error(w, common.NewNotFoundResponse([]entities.User{}).Message, common.NewBadRequestResponse([]entities.User{}).Code)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)

				json.NewEncoder(w).Encode(common.NewSuccessOperationResponse(res))

				// json.NewEncoder(w).Encode(map[string]interface{}{
				// 	"code":    200,
				// 	"message": "Succesful Operation",
				// 	"data":    res,
				// })
			}
		default:
			http.Error(w, common.NewBadRequestResponse([]entities.User{}).Message, common.NewBadRequestResponse([]entities.User{}).Code)
		}

	}
}
