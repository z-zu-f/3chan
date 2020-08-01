package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"

	"log"

	"os"

	"time"
)

var (
	dbx *sqlx.DB
	err error
)

// Res Res
type Res struct {
	RId       int       `db:"r_id"`
	ViewName  string    `db:"view_name"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}

// Thread Thread
type Thread struct {
	TId       int       `db:"t_id"`
	TTitle    string    `db:"t_title"`
	CreatedAt time.Time `db:"created_at"`
}

// Init Init
func Init() {
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = "root"
	}
	dbname := os.Getenv("MYSQL_DBNAME")
	if dbname == "" {
		dbname = "3chan"
	}
	password := os.Getenv("MYSQL_PASS")
	if password == "" {
		password = ""
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	dbx, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to DB: %s.", err.Error())
	}

	dbx.MustExec(`CREATE TABLE IF NOT EXISTS boards (r_id int AUTO_INCREMENT PRIMARY KEY, 
							t_id int, view_name varchar(25), body text, created_at datetime DEFAULT CURRENT_TIMESTAMP)
							ENGINE=InnoDB DEFAULT CHARSET=utf8;`)

	dbx.MustExec(`CREATE TABLE IF NOT EXISTS threads (t_id int AUTO_INCREMENT PRIMARY KEY, 
							t_title varchar(25), created_at datetime DEFAULT CURRENT_TIMESTAMP)
							ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
}

// FindThread FindThread
func FindThread(rID string) []Res {
	rows, err := dbx.Queryx("SELECT r_id, view_name, body, created_at FROM boards where t_id = ?", rID)
	if err != nil {
		log.Fatal(err)
	}

	var res Res
	var reslist []Res
	for rows.Next() {
		err := rows.StructScan(&res)
		if err != nil {
			log.Fatal(err)
		}
		reslist = append(reslist, res)
	}
	return reslist
}

// FindALLThreads FindALLThreads
func FindALLThreads() []Thread {
	rows, err := dbx.Queryx("SELECT t_id, t_title, created_at FROM threads")
	if err != nil {
		log.Fatal(err)
	}

	var thread Thread
	var threads []Thread
	for rows.Next() {
		err := rows.StructScan(&thread)
		if err != nil {
			log.Fatal(err)
		}
		threads = append(threads, thread)
	}
	return threads
}

// InsertThread InsertThread
func InsertThread(tTitle string) {
	_, err := dbx.Queryx("INSERT INTO threads(t_title) VALUES (?)", tTitle)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertMessage InsertMessage
func InsertMessage(tID string, viewName string, message string) {
	_, err := dbx.Queryx("INSERT INTO boards(t_id, view_name, body) VALUES (?, ?, ?)", tID, viewName, message)
	if err != nil {
		log.Printf("ddd")
		log.Fatal(err)
	}
}
