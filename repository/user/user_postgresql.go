package userRepository

import (
	"database/sql"
	"github.com/hpatel811/JWT_With_GO/models"
	"log"
)

type UserRepository struct{

}

func (u UserRepository) Signup(db *sql.DB, user models.User) (models.User,error){
	dbQuery:= "insert into users(email,password) values($1,$2) RETURNING id;"
	err:= db.QueryRow(dbQuery,user.Email,user.Password).Scan(&user.ID)
	if err!=nil{
		log.Fatal(err)
		return user,err
	}
	user.Password=""
	return user,nil
}
func (u UserRepository) Login(db *sql.DB, user models.User) (models.User,error){
	dbQuery:="select * from users where email = $1"
	row:= db.QueryRow(dbQuery,user.Email)
	err:= row.Scan(&user.ID,&user.Email,&user.Password)
	if err!=nil{
		log.Fatal(err)
		return user,err
	}
	return user,nil
}

