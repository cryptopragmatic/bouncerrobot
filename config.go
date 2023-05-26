package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct holds all our configuration
type Config struct {
	TelegramAPIKey string `json:"telegram_api_key"`
	Dev            bool   `json:"dev"`
	LogAccount     int64  `json:"log_account"`
	Votes          int64  `json:"votes"`
	Dels           int    `json:"dels"`
}

// Load method loads configuration file to Config struct
func (c *Config) load(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		log.Println(err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		log.Println(err)
	}
}

func initConfig() *Config {
	c := &Config{}
	c.load("config.json")
	return c
}
