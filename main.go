package main

import (
	"os"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router  			*mux.Router
	bindPort			string
}

func tryGetEnv(varName string, fallbackValue string) (varValue string) {
	if value, ok := os.LookupEnv(varName); ok {
		return value
	}
	return fallbackValue
}

func (a *App) InitializeRouting() {
	a.Router = mux.NewRouter()
	a.Router.StrictSlash(true)

	a.Router.HandleFunc("/status/{statusCode:[0-9]+}/", a.giveStatusCode).Methods("GET")
	a.Router.HandleFunc("/slow/", a.slowResponse).Methods("GET")

	log.Print("[INIT] Initialization complete")
}

func (a *App) Run(addr string) {

	log.Printf("Starting HTTP server at %v", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {

	a := App{
		bindPort: tryGetEnv("HTTP_PORT", "9090"),
	}

	a.InitializeRouting()

	a.Run(fmt.Sprintf("0.0.0.0:%s", a.bindPort))
}
