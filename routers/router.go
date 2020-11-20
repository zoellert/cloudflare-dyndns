package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoellert/cloudflare-dyndns/handlers"
	"log"
	"os"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	updatePassword := os.Getenv("UPDATE_PASSWORD")
	if len(updatePassword) < 1 {
		log.Fatalf("Environment variable UPDATE_PASSWORD is empty")
	}

	r.GET("/nic/update", gin.BasicAuth(gin.Accounts{
		"update": updatePassword,
	}), handlers.UpdateDNSRecord)

	return r
}
