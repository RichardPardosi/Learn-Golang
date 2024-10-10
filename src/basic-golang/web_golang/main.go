package main

import (
	"net/http"

	"github.com/RichardPardosi/try-web-crud-golang/src/basic-golang/web_golang/database"
	"github.com/RichardPardosi/try-web-crud-golang/src/basic-golang/web_golang/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
