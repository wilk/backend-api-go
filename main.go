package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/Sirupsen/logrus"
	"os"
	"net/http"
	"encoding/json"
	"github.com/martini-contrib/render"
)

func getUsers() {

}

func insertUser(res http.ResponseWriter, req *http.Request, db *gorm.DB, render render.Render) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var user User
	err := decoder.Decode(&user)
	if err != nil {
		render.HTML(http.StatusBadRequest, "Error", "Bad JSON encoding")
		return
	}

	db.Create(&user)
	db.Save(&user)

	render.JSON(http.StatusCreated, user)
}

func getUser() {

}

func updateUser() {

}

func deleteUser() {

}

type User struct {
	gorm.Model

	Name string `json:"name"{gorm:"default:'name'"`
	Age int `json:"age";gorm:"default:'age'"`
	Email string `json:"email";gorm:"default:'email'"`
	Mobile string `json:"mobile";gorm:"default:'mobile'"`
}

func main() {
	log.Info("Starting the web server...")

	log.Info("Opening the SQLite DB...")

	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer db.Close()

	log.Info("SQLite DB opened!")

	db.AutoMigrate(&User{})

	server := martini.Classic()

	server.Map(db)
	server.Use(render.Renderer())

	server.Group("/api/users", func (r martini.Router) {
		r.Get("/", getUsers)
		r.Post("/", insertUser)
		r.Get("/:id", getUser)
		r.Put("/:id", updateUser)
		r.Delete("/:id", deleteUser)
	})

	log.Info("Server listening on port 3000")

	server.Run()
}