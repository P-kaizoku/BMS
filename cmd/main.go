package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/P-kaizoku/BMS/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)

	fmt.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
