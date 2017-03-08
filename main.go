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

func getUsers(db *gorm.DB, render render.Render) {
	var users []User
	db.Find(&users)

	render.JSON(http.StatusOK, users)
}

func insertUser(req *http.Request, db *gorm.DB, render render.Render) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var user User
	err := decoder.Decode(&user)
	if err != nil {
		log.Error(err)
		render.Text(http.StatusBadRequest, "Bad JSON encoding")
		return
	}

	db.Create(&user)
	db.Save(&user)

	render.JSON(http.StatusCreated, user)
}

func getUser(params martini.Params, db *gorm.DB, render render.Render) {
	var user User
	if db.First(&user, params["id"]).RecordNotFound() {
		render.Text(http.StatusNotFound, "User not found")
		return
	}

	render.JSON(http.StatusOK, user)
}

func updateUser(req *http.Request, params martini.Params, db *gorm.DB, render render.Render) {
	var user User
	if db.First(&user, params["id"]).RecordNotFound() {
		render.Text(http.StatusNotFound, "User not found")
		return
	}

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	var userData User
	err := decoder.Decode(&userData)
	if err != nil {
		log.Error(err)
		render.Text(http.StatusBadRequest, "Bad JSON encoding")
		return
	}

	if len(userData.Name) > 0 {
		user.Name = userData.Name
	}

	if userData.Age > 0 {
		user.Age = userData.Age
	}

	if len(userData.Email) > 0 {
		user.Email = userData.Email
	}

	if len(userData.Mobile) > 0 {
		user.Mobile = userData.Mobile
	}

	db.Save(&user)

	render.JSON(http.StatusOK, user)
}

func deleteUser() {

}

type User struct {
	gorm.Model

	Name string `json:"name",omitempty;gorm:"default:'name'"`
	Age int `json:"age",omitempty;gorm:"default:'age'"`
	Email string `json:"email",omitempty;gorm:"default:'email'"`
	Mobile string `json:"mobile",omitempty;gorm:"default:'mobile'"`
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