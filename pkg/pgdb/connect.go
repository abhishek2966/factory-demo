package pgdb

import (
	"database/sql"
	"fmt"

	"github.com/abhishek2966/factory-demo/pkg/config"
	_ "github.com/lib/pq"
)

type Source struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewDBSource(base string) Source {
	return Source{
		Host:     config.ReadConfig(base, "host"),
		Port:     config.ReadConfig(base, "port"),
		User:     config.ReadConfig(base, "user"),
		Password: config.ReadConfig(base, "password"),
		DbName:   config.ReadConfig(base, "dbName"),
	}
}
func OpenConnection(s Source) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", s.Host, s.Port, s.User, s.Password, s.DbName)
	db, err := sql.Open("postgres", dataSourceName)
	return db, err
}
