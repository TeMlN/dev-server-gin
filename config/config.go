package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

//db config
var Conf Config

type Config struct {
	Database Database `toml:"database"`
}

type Database struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Server   string `toml:"server"`
	Port     string `toml:"port"`
	Engine   string `toml:"engine"`
	Database string `toml:"database"`
}

func InitConfig() {
	_, err := toml.DecodeFile("./config.toml", &Conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(Conf)

}
