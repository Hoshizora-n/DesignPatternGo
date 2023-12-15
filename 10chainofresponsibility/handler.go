package chainofresponsibility

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqId := uuid.NewString()
		log.Println(reqId)
		ctx.Set("req-id", reqId)
		ctx.Next()
	}
}

func AuthHandler(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqId := ctx.GetString("req-id")
		log.Print("auth: get req-id from logger middleware: ")
		log.Println(reqId)

		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			log.Println("auth is empty")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"success": false,
			})
			return
		}
		ctx.Next()
	}
}

func FinalHandler(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqId := ctx.GetString("req-id")
		log.Print("final: get req-id from logger middleware: ")
		log.Println(reqId)

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
		})
	}
}
