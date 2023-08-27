package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/cskarthik9/go-book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server !")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
