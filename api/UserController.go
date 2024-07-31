package api

import (
	"encoding/json"
	. "jakisRest/database"
	"log"
	"net/http"
)

type singleStringUserBody struct {
	NickName string `json:"nickname"`
}

func handleUsers(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handleGetAllUsers(writer, request)
	case http.MethodPost:
		handlePostUser(writer, request)
	case http.MethodDelete:
		handleDeleteUser(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleDeleteUser(writer http.ResponseWriter, request *http.Request) {
	var nickName singleStringUserBody
	err := json.NewDecoder(request.Body).Decode(&nickName)
	if err != nil {
		panic(err)
	}
	err = DeleteUser(nickName.NickName)
	if err != nil {
		panic(err)
	} else {
		writer.Write([]byte("Successfully deleted user " + nickName.NickName))
	}
}

func handlePostUser(writer http.ResponseWriter, request *http.Request) {
	var newUser CatUser
	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		panic(err)
	}
	err = AddUser(&newUser)
	if err != nil {
		panic(err)
	} else {
		writer.Write([]byte("User added with name: " + newUser.NickName))
	}
}

func handleGetAllUsers(writer http.ResponseWriter, request *http.Request) {
	users := GetAllUsers()

	for i := range users {
		userJson, err := json.MarshalIndent(users[i], "", " ")
		if err != nil {
			log.Fatal(err)
		} else {
			_, err := writer.Write(userJson)
			if err != nil {
				panic("error writing all users to client")
			} else {
				writer.Write([]byte("\n"))
			}
		}
	}
}
