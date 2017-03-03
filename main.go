package main

import (
	"github.com/go-martini/martini"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/Sirupsen/logrus"
	"os"
)

func getUsers() {

}

func insertUser() {

}

func getUser() {

}

func updateUser() {

}

func deleteUser() {

}

func main() {
	log.Info("Starting the web server...")

	log.Info("Opening the SQLite DB...")

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer db.Close()

	log.Info("SQLite DB opened!")

	m := martini.Classic()

	m.Group("/api/users", func (r martini.Router) {
		r.Get("/", getUsers)
		r.Post("/", insertUser)
		r.Get("/:id", getUser)
		r.Put("/:id", updateUser)
		r.Delete("/:id", deleteUser)
	})

	m.Run()

	log.Info("Server is listening on port 3000")
}