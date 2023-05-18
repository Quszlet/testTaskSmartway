package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	employeesTable = "employees"
	departmentsTable = "departments"
	companiesTable = "companies"
	passportsTable = "passports"
)

type Config struct {
	Host     string
	Port     string
	Username string
	password string
	DBName   string
	SSLmode  string
}

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
				config.Host, config.Port, config.Username, config.DBName, config.password, config.SSLmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
