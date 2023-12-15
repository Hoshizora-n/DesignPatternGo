package factory

import (
	"fmt"
	"log"
)

type Database interface {
	Connect() error
}

type MySQLDB struct{}

func (m *MySQLDB) Connect() error {
	fmt.Println("Connected to MySQL database")
	return nil
}

type PostgresDB struct{}

func (p *PostgresDB) Connect() error {
	fmt.Println("Connected to PostgreSQL database")
	return nil
}

type DatabaseFactory struct{}

func (f *DatabaseFactory) CreateDatabase(dbType string) (Database, error) {
	switch dbType {
	case "mysql":
		return &MySQLDB{}, nil
	case "postgres":
		return &PostgresDB{}, nil
	default:
		return nil, fmt.Errorf("unsupported database type")
	}
}

func Run() {
	factory := &DatabaseFactory{}

	db, err := factory.CreateDatabase("mysql")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
