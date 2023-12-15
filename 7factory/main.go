package factory

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	Connect() error
}

type MySQLDB struct {
	URL string
}

func (m *MySQLDB) Connect() error {
	fmt.Println("Connected to MySQL database")
	return nil
}

type PostgresDB struct {
	DataSource string
}

func (p *PostgresDB) Connect() error {
	fmt.Println("Connected to PostgreSQL database")
	return nil
}

type DatabaseFactory struct{}

func (f *DatabaseFactory) CreateDatabase(dbType string) (*sqlx.DB, error) {
	switch dbType {
	case "mysql":
		mysql := &MySQLDB{}
		db, _ := sqlx.Connect("mysql", mysql.URL)
		return db, nil
	case "postgres":
		postgres := &PostgresDB{}
		db, _ := sqlx.Connect("postgres", postgres.DataSource)
		return db, nil
	default:
		return nil, fmt.Errorf("unsupported database type")
	}
}

func Run() {
	factory := &DatabaseFactory{}

	_, err := factory.CreateDatabase("mysql")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database berhasil dibuat")
}
