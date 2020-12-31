package utilities

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpatel811/JWT_With_GO/models"
	"golang.org/x/crypto/bcrypt"
)
/*
func init(){
	gotenv.Load()
}*/

func ErrorResponse(w http.ResponseWriter, status int, error models.Error){
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func ResponseJSON(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(data)
}

func ComparePassword(hashedPwd string, pwd []byte) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd),pwd)
	if err!=nil{
		log.Fatal(err)
		return false
	}
	return true
}
func GenerateToken(user models.User)(string,error){
	//var err Error
	secret:=os.Getenv("SECRET")

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":user.Email,
		"iss":"course",
	})
	tokenString,err:=token.SignedString([]byte(secret))

	//spew.Dump(token)
	return tokenString,err
}
