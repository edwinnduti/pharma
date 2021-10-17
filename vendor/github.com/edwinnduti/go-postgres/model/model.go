package model

// create models
type User struct {
	ID		 uint64		`gorm:"primary_key;auto_increment" json:"id"`
	Name		string		`json:"name"`
	Age		uint64		`json:"age"`
	Location	string		`json:"location"`
}

// response struct
type Response struct {
        Code		uint64		`json:"code"`
        Message		string		`json:"message"`
}

// describe config model
type Config struct {
	Host       string
        Dbport     string
        Dbusername string
        Dbname     string
        Passwd     string
}
