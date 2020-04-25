package controllers

import (
	"fmt"
	"log"
	"net/http"

	"example.com/me/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

//Server app server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}


//Initialize initialize database server
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)
	server.DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

//Run app server
func (server *Server) Run(addr string) {
	fmt.Println("List to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
