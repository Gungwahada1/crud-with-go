package routes

import (
	"crud-with-go/controllers/DashboardController"
	"crud-with-go/controllers/UserController"
	
	"net/http"
	// "crud-with-go/controllers"
)

func InitRoutes() {
	http.HandleFunc("/", DashboardController.Index)
	http.HandleFunc("/dashboard", DashboardController.Index)

	http.HandleFunc("/users", UserController.Index)
	http.HandleFunc("/user/create", UserController.Create)
	http.HandleFunc("/user/edit", UserController.Edit)
	http.HandleFunc("/user/delete", UserController.Delete)

	// http.HandleFunc("/user/add", UserController.Add)
	// http.HandleFunc("/user/{id}/update", UserController.Update)
	// http.HandleFunc("/user/{id}/destroy", UserController.Destroy)
	
	// http.HandleFunc("/users", nil)
	// Tambahkan pengaturan rute lain di sini
}
