package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Title    string
	Host     string `toml:"host"`
	Port     int64  `toml:"port"`
	LogLevel string `toml:"log_level"`
	Server   string `toml:"Server"`
	Database string `toml:"Database"`
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
