package routes

import (
	"net/http"

	"github.com/RichardPardosi/try-web-crud-golang/src/basic-golang/web_golang/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee())
}
