package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Title    string
	Host     string `toml:"host"`
	Port     int  `toml:"port"`
	LogLevel string `toml:"log_level"`
	DBServer string `toml:"dbServer"`
	Database string `toml:"database"`
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
