package controllers

import (
	"fmt"
	"net/http"
	models "../models"
)

// UserController exporting
func UserController(rw http.ResponseWriter, req *http.Request) {
	// responseData := make([]map[string]interface{}, 0)
	// get req params
	// params := mux.Vars(req)
	// fromID, err := strconv.Atoi(params["userId"])

	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("foo controller")


	fromID := 132969

	user := models.User{}

	// userResult, err := models.GetUsers(fromID)
	userResult := user.GetUsers(fromID)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("foo ctrl -> ", userResult)

	// return userResult
	return
}
