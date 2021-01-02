package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hpatel811/JWT_With_GO/controllers"
	"github.com/hpatel811/JWT_With_GO/driver"
	"github.com/subosito/gotenv"
)

var db *sql.DB
func init(){
	gotenv.Load()
}
func main(){

	db=driver.ConnectDB()
	controller:=controllers.Controller{}
	router:=mux.NewRouter()
	router.HandleFunc("/signup",controller.Signup(db)).Methods("Post")
	router.HandleFunc("/login",controller.Login(db)).Methods("Post")
	router.HandleFunc("/protected",controller.TokenVerifyMiddleWare(controller.ProtectedEndPoint())).Methods("Get")

	//log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000",router))
}





