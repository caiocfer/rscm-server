package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	DBConnection = ""
	Port         = 0
)

func LoadConfiguration() {
	var error error

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))

	if error != nil {
		Port = 5001
	}

	DBConnection = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PROTOCOL"),
		os.Getenv("DB_NAME"),
	)
}
