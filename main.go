/*
[*] Copyright © 2021
[*] Dev/Author -> Edwin Nduti
*/

package main

// libraries to use
import (
	"log"
	"net/http"
	"os"

	"github.com/edwinnduti/pharma/router"
	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

// Main function
func main() {

	//Register router
	r := router.Router()

	//Get port
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8081"
	}

	// set CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// establish logger
	n := negroni.Classic()
	n.UseHandler(r)
	server := &http.Server{
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(n),
		Addr:    ":" + Port,
	}

	log.Printf("Listening on PORT: %s", Port)
	server.ListenAndServe()
}
