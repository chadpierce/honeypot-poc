package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	LOGS Logs	 `json:"logs"`
	TCP TCP      `json:"tcp"`
	//UDP UDP		 `json:"udp"`
}

type TCP struct {
	Ports []string `json:"ports"`
}

// type UDP struct {
// 	Ports []string `json:"ports"`
// }

type Logs struct {
	Enabled bool `json:"enabled"`
	Detailed bool `json:"detail"`
}

func Read() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("ERROR open config: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("ERROR decode config: %v", err)
	}
	return config
}
