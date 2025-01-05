package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/flamme97/ecomgo/cmd/api"
	"github.com/flamme97/ecomgo/config"
	"github.com/flamme97/ecomgo/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUSer,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	initStorage(db)
	
	router := api.GetRouter()

	http.Handle("/", router)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	 }

}


func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to DB")
}