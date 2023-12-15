package builder

type DBConnection struct {
	Provider string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Options  []string
}

func Run() {
	dbBuilder := NewDBConnectionBuilder().
		WithProvider("postgres").
		WithUsername("admin").
		WithPassword("admin").
		WithDatabase("catalog").
		AddOptions("sslmode=disable")

	db, err := dbBuilder.Build()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
