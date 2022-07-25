package database

import (
	"encoding/json"

	"github.com/nicolube/vcp-hepsiau-backend/config"
	"github.com/nicolube/vcp-hepsiau-backend/model"
)

type Reposetory interface {
	GetUser(id int) (model.UserModel, error)
	UpdateUser(user model.UserModel) error
	CreateUser(user model.UserModel) error
	DeleteUser(user model.UserModel) error

	GetTokenByUser(userId int) ([]model.UserTokenModel, error)
	GetTokenByToken(token string) (model.UserTokenModel, error)
	CreateToken(model.UserTokenModel) error
	DeleteToken(token model.UserTokenModel) error
}

type Database struct {
	Reposetories map[string]Reposetory
}

func (db *Database) Create(conf config.AppConfig) {
	db.Reposetories = make(map[string]Reposetory)
	for _, repoConfig := range conf.Reposetories {
		switch repoConfig.Type {
		case "MySQL":
		case "MariaDB":
			var sqlConf config.SQLConfig
			if err := json.Unmarshal(repoConfig.DataRaw, &sqlConf); err != nil {
				panic(err)
			}
			db.Reposetories[repoConfig.Name] = conntectToSql(sqlConf, "mysql")
		}

	}
}
