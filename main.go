package main

import (
	"rscm/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfiguration()

}
