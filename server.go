package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nicolube/vcp-hepsiau-backend/config"
	"github.com/nicolube/vcp-hepsiau-backend/database"
	"github.com/nicolube/vcp-hepsiau-backend/router"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Load config
	dat, err := os.ReadFile("config.json")
	check(err)
	conf := config.LoadConfig(string(dat))

	// Start Database
	db := database.Database{}
	db.Create(conf)

	// Setup router
	rtr := mux.NewRouter()
	router.Init(rtr, db.Reposetories["default"])
	http.Handle("/", rtr)
	// Spinup Server
	host := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	fmt.Printf("Server is listening on http://%s\n", host)
	fmt.Println(http.ListenAndServe(host, handlers.CORS()(rtr)))
}
