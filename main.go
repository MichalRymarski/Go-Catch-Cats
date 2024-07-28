package main

import (
	. "jakisRest/api"
	. "jakisRest/database"
)

func main() {
	SetupDatabase()
	SetupRestApi()
}
