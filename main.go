package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rscm/src/config"
	routes "rscm/src/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfiguration()

	fmt.Printf("Service started.\nListening to port: %d.\n", config.Port)

	r := routes.GenerateRoutes()
	port := ":" + os.Getenv("API_PORT")

	log.Fatal(http.ListenAndServe(port, r))

}
