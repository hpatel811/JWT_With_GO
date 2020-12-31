package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpatel811/JWT_With_GO/models"
	"github.com/hpatel811/JWT_With_GO/repository/user"
	"github.com/hpatel811/JWT_With_GO/utilities"
	"golang.org/x/crypto/bcrypt"
)
type Controller struct{

}
//Signup endpoint handled creating a user
func (c Controller) Signup(db *sql.DB) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){
		var user models.User
		var error models.Error
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email==""||user.Password==""{
			error.Message ="Email and/or Password is missing in the request."
			utilities.ErrorResponse(w,http.StatusBadRequest,error)
			return
		}
		hash,err:= bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
		if err!=nil {
			log.Fatal(err)
		}
		user.Password = string(hash)
		user,err=userRepository.UserRepository{}.Signup(db,user)
		if err!=nil{
			error.Message ="Server Error."
			utilities.ErrorResponse(w,http.StatusInternalServerError,error)
			return
		}
		//responding with empty password, so we dont expose the password
		user.Password=""
		w.Header().Set("Content-Type","application/JSON")
		utilities.ResponseJSON(w,user)
		return
		//spew.Dump(user)
	}
}
//Login endpoint handled verifying the user and generating a token for the user's login
func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var user models.User
		var jwt models.JWT
		var error models.Error
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email==""||user.Password==""{
			error.Message ="Email and/or Password is missing in the request."
			utilities.ErrorResponse(w,http.StatusBadRequest,error)
			return
		}
		password:= user.Password
		user,err:=userRepository.UserRepository{}.Login(db,user)
		if err!=nil{
			//log.Fatal(err)
			if err==sql.ErrNoRows{
				error.Message="The user doesnt exist"
				utilities.ErrorResponse(w,http.StatusBadRequest,error)
				return
			}else{
				log.Fatal(err)
			}
		}
		hashedPassword:=user.Password
		pwdMatch:=utilities.ComparePassword(hashedPassword,[]byte(password))
		if !pwdMatch{
			error.Message="Invalid Password"
			utilities.ErrorResponse(w,http.StatusBadRequest,error)
			return
		}
		//spew.Dump(user)

		token,err:= utilities.GenerateToken(user)
		if err!=nil{
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		jwt.Token=token
		utilities.ResponseJSON(w,jwt)
	}
}
//TokenVerifyMiddleWare will verify the token provided, so it will direct to the next protected endpoint
func (c Controller) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var errorObj models.Error
		bearerToken:= strings.Split(r.Header.Get("Authorization")," ")
		if len(bearerToken)==2{
			authToken:=bearerToken[1]
			tkn,err:=jwt.Parse(authToken,func(token *jwt.Token)(interface{}, error){
				if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
					return nil,fmt.Errorf("There was an error")
				}
				return []byte(os.Getenv("SECRET")),nil
			})
			if err!=nil{
				errorObj.Message=err.Error()
				utilities.ErrorResponse(w,http.StatusUnauthorized,errorObj)
				return
			}
			if tkn.Valid{
				next.ServeHTTP(w,r)
			}else{
				errorObj.Message=err.Error()
				utilities.ErrorResponse(w,http.StatusUnauthorized,errorObj)
				return
			}
			//spew.Dump(tkn)
		}else{
			errorObj.Message="Invalid token"
			utilities.ErrorResponse(w,http.StatusUnauthorized,errorObj)
			return
		}
	})
}


