package main

import (
	"fmt"
	"log"
	"net/http"

	"PBP-APITools/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	controller.RunScheduler()

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
