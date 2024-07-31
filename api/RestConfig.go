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
	mux.HandleFunc("/example-user", handleExampleUser)
	mux.HandleFunc("/default-admin", handleDefaultAdmin)
	mux.HandleFunc("/folders", handleFolders)
	mux.HandleFunc("/users", handleUsers)
	return mux
}

func handleDefaultAdmin(writer http.ResponseWriter, request *http.Request) {
	err := AddDefaultAdmin()
	if err != nil {
		log.Panic(err)
	}
}

func handleExampleUser(writer http.ResponseWriter, request *http.Request) {
	err := AddExampleUser()
	if err != nil {
		log.Panic(err)
	}
	writer.Write([]byte("Example user added"))
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	users := GetAllUsers()
	for i := range users {
		userJson, err := json.MarshalIndent(users[i], "", "  ")
		if err != nil {
			log.Fatal(err)
		} else {
			_, err := writer.Write(userJson)
			if err != nil {
				panic("error writing to client")
			} else {
				writer.Write([]byte("\n"))
			}
		}
	}
}
