package chainofresponsibility

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Log *log.Logger

func Run() {
	r := gin.New()

	Log = log.Default()

	r.Use(Logger(Log))
	r.Use(AuthHandler(Log))

	r.GET("/hello", FinalHandler(Log))

	// Start the server
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
