package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/Nortvest/goBase64/internal/app/gobase64"
)

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "configs/prod.toml", "path to config path")
}

func main() {
	flag.Parse()

	config := gobase64.NewConfig()
	_, err := toml.DecodeFile(ConfigPath, config)
	if err != nil {
		panic(fmt.Sprintf("[Config error] %s", err.Error()))
	}

	server := gobase64.New(config)

	if err := server.Start(); err != nil {
		panic(fmt.Sprintf("[Server error] %s", err.Error()))
	}
}
