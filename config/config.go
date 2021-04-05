package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

//db config
var Conf Config

type Config struct {
	APP      app
	Database database
	Format   format
}

type database struct {
	User     string
	Password string
	Server   string
	Port     string
	Engine   string
	Database string
}

type format struct {
	Format string
}

type app struct {
	Name string `toml:"name"`
}

func StartConfig() {
	_, err := toml.DecodeFile("./config.toml", &Conf)
	err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", Conf)

}
