package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	database database
	server   server
}

func (c config) GetDatabase() database {
	return c.database
}

func (c config) GetServer() server {
	return c.server
}

var c config

func Init() error {
	var err error

	err = godotenv.Load()

	if err != nil {
		return err
	}

	serverPort, err := strconv.ParseUint(os.Getenv("PORT"), 10, 64)

	if err != nil {
		return err
	}

	c = config{
		database: database{
			host: os.Getenv("DATABASE_HOST"),
			user: os.Getenv("DATABASE_USER"),
			pass: os.Getenv("DATABASE_PASS"),
			name: os.Getenv("DATABASE_NAME"),
			port: os.Getenv("DATABASE_PORT"),
		},
		server: server{
			port: uint(serverPort),
		},
	}

	return nil
}

func GetConfig() *config {
	return &c
}
