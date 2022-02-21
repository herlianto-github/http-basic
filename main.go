package main

import (
	"fmt"
	"httpBasic/configs"
	"httpBasic/delivery/controllers/users"
	"httpBasic/delivery/routes"
	usersRepo "httpBasic/repository/users"
	"httpBasic/utils"
	"net/http"
)

func main() {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	handlers := http.NewServeMux()

	usersRepo := usersRepo.NewUserRepo(db)
	usersCtrl := users.NewUsersControllers(usersRepo)

	routes.Endpoints(handlers, usersCtrl)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), handlers)

}
