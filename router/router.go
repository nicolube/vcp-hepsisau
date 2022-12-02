package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nicolube/vcp-hepsiau-backend/database"
	"github.com/nicolube/vcp-hepsiau-backend/model"
)

func Init(rtr *mux.Router, repo database.Reposetory) {
	userManager := UserManager{
		Reposetory: repo,
	}

	rtr.HandleFunc("/", RootHandler)
	rtr.Path("/content/{id:[0-9]+}").HandlerFunc(getContent(repo)).Methods("GET")
	rtr.Path("/menu").HandlerFunc(getMenu(repo)).Methods("GET")
	rtr.Path("/site").HandlerFunc(getSite(repo)).Methods("GET")
	securedRtr := rtr.PathPrefix("/secure").Subrouter()
	securedRtr.Use(userManager.Auth)
	securedRtr.HandleFunc("/content", createContent(repo)).Methods("PUT")
	securedRtr.HandleFunc("/menu", saveMenu(repo)).Methods("POST")
	securedRtr.Path("/content/{id:[0-9]+}").HandlerFunc(deleteContent(repo)).Methods("DELETE")
	securedRtr.NewRoute().HandlerFunc(notFound)
}

func checkErr(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return true
	}
	return false
}

func notFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
}

func createContent(repo database.Reposetory) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.Path)
		user := req.Context().Value(UserManagerContent("user")).(model.UserModel)
		var content model.ContentModel
		json.NewDecoder(req.Body).Decode(&content)
		content.UserId = user.Id
		content, err := repo.CreateContent(content)
		if checkErr(w, err) {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func deleteContent(repo database.Reposetory) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(mux.Vars(req)["id"])
		if checkErr(w, err) {
			return
		}
		content := model.ContentModel{}
		content.Id = int64(id)
		err = repo.DeleteContent(content)
		if checkErr(w, err) {
			return
		}
		w.WriteHeader(http.StatusOK)

	}
}

func getContent(repo database.Reposetory) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
		if checkErr(w, err) {
			return
		}
		content, err := repo.GetContent(id)
		if checkErr(w, err) {
			return
		}
		data, err := json.Marshal(content)
		if checkErr(w, err) {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func getMenu(repo database.Reposetory) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		menu, err := repo.GetMenu()
		if checkErr(w, err) {
			return
		}
		data, err := json.Marshal(menu)
		if checkErr(w, err) {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func saveMenu(repo database.Reposetory) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var menu []model.MenuItemModel
		err := json.NewDecoder(req.Body).Decode(&menu)
		if checkErr(w, err) {
			return
		}
		err = repo.SaveMenu(menu)
		if checkErr(w, err) {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func getSite(repo database.Reposetory) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		url := req.URL.Query().Get("url")
		if len(url) == 0 {
			checkErr(w, fmt.Errorf("no url is given"))
			return
		}
		menu, err := repo.GetSite(url)
		if checkErr(w, err) {
			return
		}
		data, err := json.Marshal(menu)
		if checkErr(w, err) {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
