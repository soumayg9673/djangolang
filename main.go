package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	serverCtrl "github.com/soumayg9673/djangolang/djangolang"
)

func init() {
	log.Println("DjanGoLang")

	serverCtrl.ReadEnv()
	serverCtrl.ConnectDb()
}

func main() {
	serverCtrl.BeServer.RTR = mux.NewRouter()
	log.Println("Initializing HTTP server on :8080")
	err := http.ListenAndServe(":8080", serverCtrl.BeServer.RTR)
	if err != nil {
		log.Fatalf("Error starting in server. Error %v", err)
	}

	serverCtrl.BeServer.DB.Close()
}
