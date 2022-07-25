package config

import (
	"encoding/json"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type AppConfig struct {
	Host         string              `json:"host"`
	Port         int                 `json:"port"`
	Reposetories []RepositoriyConfig `json:"repositories"`
}

type SQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"username"`
	Database string `json:"database"`
	Password string `json:"password"`
}

type RepositoriyConfig struct {
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	DataRaw json.RawMessage `json:"data"`
}

func LoadConfig(data string) AppConfig {
	var config AppConfig
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		panic(err)
	}
	return config
}

func (cfg *SQLConfig) ToSqlConfig() string {
	out := mysql.Config{
		User:                 cfg.User,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DBName:               cfg.Database,
		AllowNativePasswords: true,
	}

	return out.FormatDSN()
}
