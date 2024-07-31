package api

import (
	"encoding/json"
	. "jakisRest/database"
	"log"
	"net/http"
)

type singleStringFolderBody struct {
	Name string `json:"name"`
}

func handleFolders(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handleGetAllFolders(writer, request)
	case http.MethodPost:
		handlePostFolder(writer, request)
	case http.MethodDelete:
		handleDeleteFolder(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleDeleteFolder(writer http.ResponseWriter, request *http.Request) {
	var folderName singleStringFolderBody
	err := json.NewDecoder(request.Body).Decode(&folderName)
	if err != nil {
		panic(err)
	}
	err = DeleteFolder(folderName.Name)
	if err != nil {
		panic(err)
	}
}
func handlePostFolder(writer http.ResponseWriter, request *http.Request) {
	var folderName singleStringFolderBody
	err := json.NewDecoder(request.Body).Decode(&folderName)
	if err != nil {
		panic("failed to decode folder from post method")
	}
	err = AddFolder(folderName.Name)
	if err != nil {
		panic(err)
	} else {
		writer.Write([]byte("Folder added with name: " + folderName.Name))
	}
}

func handleGetAllFolders(writer http.ResponseWriter, request *http.Request) {
	folders := GetAllFolders()
	for i := range folders {
		folderJson, err := json.MarshalIndent(folders[i], "", "  ")
		if err != nil {
			log.Fatal(err)
		} else {
			_, err := writer.Write(folderJson)
			if err != nil {
				panic("error writing all folders to client")
			} else {
				writer.Write([]byte("\n"))
			}
		}
	}
}
