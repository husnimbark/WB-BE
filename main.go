package main

import (
	"fmt"
	"net/http"
	"waysbucks/database"
	"waysbucks/pkg/mysql"
	"waysbucks/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	fmt.Println("server running localhost:4000")
	http.ListenAndServe("localhost:4000", r)

	// env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
}
