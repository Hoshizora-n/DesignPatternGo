package builder

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConnectionBuilder struct {
	connection DBConnection
}

// Default value, port: 5432, host: localhost
func NewDBConnectionBuilder() *DBConnectionBuilder {
	return &DBConnectionBuilder{connection: DBConnection{Port: 5432, Host: "localhost"}} // Default port
}

func (builder *DBConnectionBuilder) WithProvider(p string) *DBConnectionBuilder {
	builder.connection.Provider = p
	return builder
}

func (builder *DBConnectionBuilder) WithHost(host string) *DBConnectionBuilder {
	builder.connection.Host = host
	return builder
}

func (builder *DBConnectionBuilder) WithPort(port int) *DBConnectionBuilder {
	builder.connection.Port = port
	return builder
}

func (builder *DBConnectionBuilder) WithUsername(username string) *DBConnectionBuilder {
	builder.connection.Username = username
	return builder
}

func (builder *DBConnectionBuilder) WithPassword(password string) *DBConnectionBuilder {
	builder.connection.Password = password
	return builder
}

func (builder *DBConnectionBuilder) WithDatabase(database string) *DBConnectionBuilder {
	builder.connection.Database = database
	return builder
}

func (builder *DBConnectionBuilder) AddOptions(option string) *DBConnectionBuilder {
	builder.connection.Options = append(builder.connection.Options, option)
	return builder
}

func (builder *DBConnectionBuilder) Build() (*sqlx.DB, error) {
	url := (&url.URL{
		Scheme:   builder.connection.Provider,
		User:     url.UserPassword(builder.connection.Username, builder.connection.Password),
		Host:     fmt.Sprintf("%s:%d", builder.connection.Host, builder.connection.Port),
		Path:     builder.connection.Database,
		RawQuery: strings.Join(builder.connection.Options, "&"),
	}).String()

	db, err := sqlx.Connect(builder.connection.Provider, url)
	if err != nil {
		return nil, err
	}

	fmt.Println("successfully connect to db")
	return db, nil
}
