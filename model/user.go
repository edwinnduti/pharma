package model

// create models
type User struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	NationalId string `json:"nationalid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Dob        string `json:"dob"`
	//	Age        uint64 `json:"age"`
	Location string `json:"location"`
}

// response struct
type Response struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}

// describe config model
type Config struct {
	Host       string
	Dbport     string
	Dbusername string
	Dbname     string
	Passwd     string
}
