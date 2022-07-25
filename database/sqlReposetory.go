package database

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/nicolube/vcp-hepsiau-backend/config"
	"github.com/nicolube/vcp-hepsiau-backend/model"
)

type SQLReposetory struct {
	db *sql.DB
	Reposetory
}

func (repo *SQLReposetory) runScript(path string) {
	log.Println("Run: " + path)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	querrys := strings.Split(string(data), ";")
	for _, q := range querrys {
		if len(q) == 0 {
			continue
		}
		_, err := repo.db.Exec(q + ";")
		if err != nil {
			panic(err)
		}
	}
}

func conntectToSql(conf config.SQLConfig, driver string) *SQLReposetory {
	dbUrl := conf.ToSqlConfig()
	log.Printf("Connect to %s\n", dbUrl)
	repo := new(SQLReposetory)
	var err error
	if repo.db, err = sql.Open(driver, dbUrl); err != nil {
		log.Fatal(err)
		panic(err)
	}
	repo.db.SetConnMaxLifetime(0)
	repo.db.SetMaxIdleConns(3)
	repo.db.SetMaxOpenConns(3)
	repo.runScript("database/schema/tables.sql")
	return repo
}

func (repo *SQLReposetory) GetUser(id int) (model.UserModel, error) {
	model := model.UserModel{}
	row := repo.db.QueryRow("SELECT * FROM user WHERE id=?", id)
	err := row.Scan(&model.Id,
		&model.FirstName,
		&model.LastName, &model.Email, &model.Password, &model.LastLoginAt, &model.SingedUpAt)
	return model, err
}

func (repo *SQLReposetory) UpdateUser(user model.UserModel) error {
	_, err := repo.db.Exec("UPDATE user first_name = ?, last_name = ?, email = ?, password = ?, lastLogin = ? WHERE id=?",
		user.FirstName, user.LastName, user.Email, user.Password, user.LastLoginAt, user.Id)
	return err
}

func (repo *SQLReposetory) CreateUser(user model.UserModel) error {
	_, err := repo.db.Exec("INSERT INTO user (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password, user.LastLoginAt)
	return err
}

func (repo *SQLReposetory) DeleteUser(user model.UserModel) error {
	_, err := repo.db.Exec("DELETE FROM user WHERE id=?", user.Id)
	return err
}

func (repo *SQLReposetory) GetTokenByUser(userId int) ([]model.UserTokenModel, error) {
	tokens := make([]model.UserTokenModel, 0)
	rows, err := repo.db.Query("SELECT * FROM user_token WHERE id=?", userId)
	if err != nil {
		return tokens, err
	}
	defer rows.Close()
	for rows.Next() {
		token := model.UserTokenModel{}
		err = rows.Scan(&token.Id, &token.UserId, &token.Token, &token.Ip, &token.LastUsedAt, &token.CreatedAt)
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	return tokens, err
}

func (repo *SQLReposetory) GetTokenByToken(token string) (model.UserTokenModel, error) {
	tokenModel := model.UserTokenModel{}
	row := repo.db.QueryRow("SELECT * FROM user_token WHERE token=?", token)
	err := row.Scan(&tokenModel.Id, &tokenModel.UserId, &tokenModel.Token, &tokenModel.Ip, &tokenModel.LastUsedAt, &tokenModel.CreatedAt)
	return tokenModel, err
}

func (repo *SQLReposetory) CreateToken(token model.UserTokenModel) error {
	_, err := repo.db.Exec("INSERT INTO user_token (user_id, ip, email) VALUES (?, ?, ?)",
		token.UserId, token.Token, token.Ip)
	return err
}

func (repo *SQLReposetory) DeleteToken(token model.UserTokenModel) error {
	_, err := repo.db.Exec("DELETE FROM user_token WHERE id=?", token.Id)
	return err
}
