package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var db *sql.DB
func init(){
	gotenv.Load()
}
func ConnectDB()*sql.DB{
	//We are parsing the URL here and fetching the DB details using PQ.
	pgUrl, err := pq.ParseURL(os.Getenv("SQL_URL"))
	if err!=nil {
		log.Fatal(err)
	}

	//SQL open. this opens a connection for us, to perform any DB operations.
	db,err = sql.Open("postgres",pgUrl)
	if err!=nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err!=nil{
		log.Fatal(err)
	}
	return db
}
