package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/middleware"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	ToppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(ToppingRepository)

	r.HandleFunc("/toppings", h.FindToppings).Methods("GET")
	r.HandleFunc("/topping/{id}", h.GetTopping).Methods("GET")
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.CreateTopping))).Methods("POST")
}
