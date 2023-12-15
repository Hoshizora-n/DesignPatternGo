package singleton

import (
	"design-pattern/6singleton/config"
	"fmt"
)

func Run() {
	cfg := config.GetConfig()
	cfg2 := config.GetConfig()

	fmt.Println(cfg == cfg2)

	address := fmt.Sprintf(":%d", cfg.Server.Port)
	fmt.Printf("server run on port %s\n", address)
	// http.ListenAndServe(address, nil)
}
