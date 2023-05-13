package Config

import (
	"database/sql"
	"fmt"
	"time"
)

type Credentials struct {
	Username string
	Password string
	Server   string
	Dbname   string
}

var Database = Credentials{
	Username: "root",
	Password: "Ranjith",
	Server:   "tcp(localhost:3306)",
	Dbname:   "sachindb",
}

// ConnectToDB connects to the database --> orders
func (m Credentials) ConnectToDB() *sql.DB {
	dataSourceName := m.Username + ":" + m.Password + "@" + m.Server + "/" + m.Dbname
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour * 1)
	fmt.Println("Connected to DB Successfully....... ")
	return db
}

// NewTable creates new table if the table not exist
func NewTable() {
	db := Database.ConnectToDB()
	defer db.Close()
	_, err := db.Query("CREATE TABLE IF NOT EXISTS users(Name varchar(20) UNIQUE NOT NULL, Username varchar(20) NOT NULL, Email varchar(20) NOT NULL, Password varchar(20) NOT NULL)")
	if err != nil {
		fmt.Println(err)
	}
	_, e := db.Query("CREATE TABLE IF NOT EXISTS todo(ID int(5) UNIQUE NOT NULL AUTO_INCREMENT, Title varchar(20) NOT NULL, Description varchar(20))")
	if e != nil {
		fmt.Println(e)
	}
}
