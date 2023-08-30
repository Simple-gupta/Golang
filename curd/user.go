package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
const DNS="root:Sysdba_1@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"

type User struct{
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
}

func InitialMigration(){
	DB, err = gorm.Open(mysql.Open(DNS),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}  

func GetUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user []User
	DB.First(&user,params["id"])
	json.NewEncoder(w).Encode(user)

}  

func GetUserCustom(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user []User
	DB.First(&user,4)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}  

func GetUserByCustomQuery(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	DB.Raw("Select * from users where ID=?",4).Scan(&user)
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user []User
	DB.First(&user,params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}  

func DeleteUser(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user []User
	DB.Delete(&user,params["id"])
	json.NewEncoder(w).Encode("the user is deleted")
}  

