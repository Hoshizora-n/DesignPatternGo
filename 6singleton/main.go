package singleton

import (
	"design-pattern/6singleton/config"
	"fmt"
)

func Run() {
	err := config.LoadConfig("./6singleton")
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf(":%d", config.Configuration.Server.Port)
	fmt.Printf("server run on port %s\n", address)
	// http.ListenAndServe(address, nil)
}
