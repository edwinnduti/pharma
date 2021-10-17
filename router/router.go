package router

import (
	// "github.com/edwinnduti/pharma/middleware"
	"github.com/edwinnduti/pharma/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// router.HandleFunc("/", middleware.HelloHandler).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api", lib.PostDataHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api", middleware.PostDataHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{user_id}", middleware.GetUserHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users", middleware.GetAllUsersHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/{user_id}", middleware.UpdateUserHandler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/user/{id}", middleware.DeleteUserHandler).Methods("DELETE", "OPTIONS")

	return router
}
