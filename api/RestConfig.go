package api

import (
	"encoding/json"
	"errors"
	. "jakisRest/database"
	"log"
	"net/http"
)

var server *http.Server

func SetupRestApi() {
	startServer()
}

func startServer() {
	mux := setupRouter()

	server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}
}

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)

	return mux
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	users := GetAllUsers()
	for i := range users {

		userJson, err := json.Marshal(users[i])
		if err != nil {
			log.Fatal(err)
		} else {
			writer.Write(userJson)
		}
	}
}
