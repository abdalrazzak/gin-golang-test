package controllers

import "github.com/abdalrazzak/gin-golang-test/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Files routes
	s.Router.HandleFunc("/files", middlewares.SetMiddlewareJSON(s.CreateFile)).Methods("POST")
	s.Router.HandleFunc("/files", middlewares.SetMiddlewareJSON(s.GetFiles)).Methods("GET")  
	s.Router.HandleFunc("/files/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteFile)).Methods("DELETE")
}
