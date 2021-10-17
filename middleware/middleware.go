package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/edwinnduti/pharma/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// declare db,err and their types
var db *gorm.DB
var err error
var config model.Config

// init function
func init() {
	fmt.Println("Getting configs...")
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env:  %v\n", err)
		var scrub http.ResponseWriter

		scrub.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "GettingEnvError!",
		}
		// give response to client
		json.NewEncoder(scrub).Encode(response)

	}

	fmt.Println("We are getting the env values")

	// secret keys
	config.Host = os.Getenv("HOST")
	config.Dbport = os.Getenv("DBPORT")
	config.Dbusername = os.Getenv("USER")
	config.Dbname = os.Getenv("DBNAME")
	config.Passwd = os.Getenv("PASSWORD")
}

// connect to database
func ConnectDb() (*gorm.DB, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.Host, config.Dbport, config.Dbusername, config.Dbname, config.Passwd)
	db, err = gorm.Open("postgres", dbURL)
	return db, err
}

// book-in new user
func PostDataHandler(w http.ResponseWriter, r *http.Request) {
	// connect to database
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("DB Error: %v\n", err)

		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DbConnectionError!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	fmt.Println("database connected")

	// make migrations
	db.AutoMigrate(&model.User{})
	fmt.Println("database migrated")

	// create an empty user struct
	var user model.User

	// decode incoming values to user empty struct
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("updating body gave error: %v\n", err)
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DataEnteringError!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	fmt.Println("decoding achieved")

	// insert user to database
	db.Create(&user)
	fmt.Println("offender inserted")

	// response
	response := model.Response{
		Code:    200,
		Message: "rows created!",
	}

	json.NewEncoder(w).Encode(response)
}

// read data of one user <identifier>
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// get id from identifier
	identifier := mux.Vars(r)
	u64, err := strconv.ParseUint(identifier["user_id"], 10, 64)
	if err != nil {
		log.Fatalf("error parsing identifier: %v\n", err)
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "ParsingIdentifierFailed!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}
	user_id := uint(u64)

	var user model.User

	// connect to database
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("Database Connection Error: %v\n", err)

		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DataEnteringError!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	// check in db
	err = db.First(&user, user_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("Record not found in database: %v\n", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		} else {
			log.Fatalf("Record not found in database: %v\n", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		}
	}

	// give response to client
	json.NewEncoder(w).Encode(user)
}

// read data of all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// models slice
	var users []model.User

	// connect to database
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)

		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DataEnteringError!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	// fetch in db
	result := db.Find(&users)
	if result.Error != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("Record not found in database: %v\n", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		} else {
			log.Fatalf("getting data error: %v\n", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		}
	}

	// give response to client
	json.NewEncoder(w).Encode(users)
}

// update user data
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// get id from identifier
	identifier := mux.Vars(r)
	u64, err := strconv.ParseUint(identifier["user_id"], 10, 64)
	if err != nil {
		log.Fatalf("error parsing identifier: %v\n", err)
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "ParsingIdentifierFailed!",
		}

		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	user_id := uint(u64)

	// declare users
	var oldUser model.User
	var newUser model.User

	// deposit to newUser
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Fatalf("updating body gave error: %v\n", err)
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DataEnteringError!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	// connect to database
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)

		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DataEnteringError!",
		}

		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	// check in db
	err = db.First(&oldUser, user_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("Record was not found: %v", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		} else {
			log.Fatalf("Record was not found: %v\n", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		}
	}

	// update model in db
	db.Model(&oldUser).Update(newUser)

	// give response to user
	json.NewEncoder(w).Encode(newUser)
}

// delete user <identifier>
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// get id from identifier
	identifier := mux.Vars(r)
	u64, err := strconv.ParseUint(identifier["user_id"], 10, 64)
	if err != nil {
		log.Fatalf("error parsing identifier: %v\n", err)
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DataEnteringError!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	user_id := uint(u64)

	// declare user
	var user model.User

	// connect to database
	db, err := ConnectDb()
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code:    500,
			Message: "DataEnteringError!",
		}
		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	// search in database
	err = db.First(&user, user_id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("Record was not found: %v", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		} else {
			log.Fatalf("Record was not found: %v", err)
			w.WriteHeader(http.StatusOK)
			response := model.Response{
				Code:    500,
				Message: "RecordNotFound",
			}
			// give response to client
			json.NewEncoder(w).Encode(response)
		}
	}

	// delete function
	db.Delete(&user)

	// pass response to client
	response := model.Response{
		Code:    user.ID,
		Message: "rows deleted!",
	}
	json.NewEncoder(w).Encode(response)
}
