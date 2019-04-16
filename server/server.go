package main

import (
	"log"
	"net/http"
	"os"

	"jwt-tools-online-shit/config"
)

const defaultDeployMode = "localDev"

func main() {
	log.Printf("[INFO] JWT-Tools-Online API REST - starting server ...")

	deployMode := defaultDeployMode
	if os.Getenv("DEPLOY_MODE") != "" {
		deployMode = os.Getenv("DEPLOY_MODE")
	}

	c, err := config.LoadConfiguration(deployMode)
	if err != nil {
		log.Fatal(err)
	}

	port := c.GetServerPort()
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	log.Printf("[INFO] server endpoint: %s \n", c.GetAPIEndPoint())

	router := newRouter(*c)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
