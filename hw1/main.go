package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const (
	host     = "database"
	port     = 5432
	user     = "postgres"
	password = "postgrespw"
	dbname   = "helloer"
)

func hellos(db *sql.DB) func(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	return func(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
		result, err := db.Query(`select hello from hellos`)
		if err != nil {
			panic(err)
		}
		for result.Next() {
			var text string
			if err := result.Scan(&text); err != nil {
				panic(err)
			}
			if _, err := rw.Write([]byte(text + "\n")); err != nil {
				panic(err)
			}
		}
	}
}

func myHello(db *sql.DB) func(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	return func(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(`insert into hellos(hello) values ($1)`, body)
		if err != nil {
			panic(err)
		}
		if _, err := rw.Write([]byte("ok")); err != nil {
			panic(err)
		}
	}
}

func main() {
	time.Sleep(10 * time.Second)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	mux := httprouter.New()
	mux.GET("/hellos", hellos(db))
	mux.POST("/my_hello", myHello(db))
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
