package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gzim07/go-bookstore/pkg/config"
	"github.com/gzim07/go-bookstore/pkg/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading enviroment variables")
	}
	config.ConnectToDb()
}
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
