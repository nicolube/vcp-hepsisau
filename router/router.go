package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolube/vcp-hepsiau-backend/database"
)

func Init(rtr *mux.Router, repo database.Reposetory) {
	userManager := UserManager{
		Reposetory: repo,
	}

	rtr.HandleFunc("/", RootHandler)
	// securedRtr := rtr.PathPrefix("/secure").MatcherFunc(authMatcher).Subrouter()
	securedRtr := rtr.PathPrefix("/secure").Subrouter()
	securedRtr.Use(userManager.Auth)
	securedRtr.HandleFunc("/test", TestSecuredHandler).Methods("GET")
	securedRtr.HandleFunc("/test2/", TestSecured2Handler)
	securedRtr.NewRoute().HandlerFunc(notFound)
	// rtr.PathPrefix("/secure").HandlerFunc(noAuth)
}

func notFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
}

func TestSecuredHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Welcome to the secured area!\n")
	fmt.Fprintf(w, "Welcome to the secured area!")
}
func TestSecured2Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the secured area!2")
}
