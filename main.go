package main

import (
	"fmt"
	"net/http"
	"rest-api/env"
	"rest-api/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	env := env.GetEnv()

	apiPort := env["API_PORT"]

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	server := &http.Server{
		Addr:         ":" + apiPort,
		Handler:      routes.GetRoutes(),
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
	}

	fmt.Println("Starting server on port: " + apiPort)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}
