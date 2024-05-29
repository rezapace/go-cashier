package routes

import (
	"myapp/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/barang", controllers.GetAllBarang).Methods("GET")
	router.HandleFunc("/barang/{id}", controllers.GetBarangByID).Methods("GET")
	router.HandleFunc("/barang", controllers.CreateBarang).Methods("POST")
	router.HandleFunc("/barang/{id}", controllers.UpdateBarang).Methods("PUT")
	router.HandleFunc("/barang/{id}", controllers.DeleteBarang).Methods("DELETE")
}
