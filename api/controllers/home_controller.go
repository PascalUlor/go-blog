package controllers

import (
	"net/http"

	"example.com/me/responses"
)

//Home homepage controller
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To MY Go Playground")
}
