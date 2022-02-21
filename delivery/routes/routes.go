package routes

import (
	"httpBasic/delivery/controllers/users"
	"net/http"
)

func Endpoints(ht *http.ServeMux, uctrl *users.UsersController) {

	ht.Handle("/users", uctrl.Gets())

}
