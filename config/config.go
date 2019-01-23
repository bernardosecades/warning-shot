package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type mySql struct {
	Host     string `env:"MYSQL_HOST" envDefault:"localhost"`
	Port     string `env:"MYSQL_PORT" envDefault:"3306"`
	User     string `env:"MYSQL_USER" envDefault:"root"`
	Password string `env:"MYSQL_PASSWORD" envDefault:"root"`
	Database string `env:"MYSQL_DATABASE" envDefault:"myDb"`
}

type webServer struct {
	Port     string    `env:"WEBSERVER_PORT" envDefault:"3000"`
}

type config struct {
	MySql         mySql
	WebServer     webServer
}

func GetConfig() config {

	m := mySql{}
	errM := env.Parse(&m)

	if errM != nil {
	  fmt.Printf("%+v\n", errM)
	}

	w := webServer{}
	errW := env.Parse(&w)

	if errW != nil {
		fmt.Printf("%+v\n", errW)
	}

	return config{m, w}
}
