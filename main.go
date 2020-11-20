package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zoellert/cloudflare-dyndns/routers"
	"github.com/zoellert/cloudflare-dyndns/services"
	"log"
	"net/http"
	"os"
)

func init() {
	services.InitCloudflare()
}

func main() {

	gin.SetMode(os.Getenv("APP_MODE"))

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	log.Printf("Starting http server listening on %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Print("Failed to start http server")
		log.Fatal(err)
	}

}
