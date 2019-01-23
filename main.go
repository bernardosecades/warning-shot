package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"warning-shot/config"

	"github.com/gorilla/mux"
)

func myService(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	io.WriteString(w, "Hello world!"+vars["user_id"])

}

// compile passing -ldflags "-X main.Build=<build sha1>"
var Build string

func main() {

	fmt.Printf("Using build: %s\n", Build)

	cf := config.GetConfig()
	r := mux.NewRouter()

	r.HandleFunc("/warning/{user_id}", myService).Methods(http.MethodGet)

	err := http.ListenAndServe(":"+cf.WebServer.Port, r)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
