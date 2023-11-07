package main

import (
	"encoding/json"
	"log"
	"singleton-pattern/config"
)

func main() {
	config := config.GetConfig()

	configBytes, err := json.Marshal(config)
	if err != nil {
		log.Println("Error: " + err.Error())
		log.Fatalln("Exiting Program on failure...")
	}

	log.Println("Successfully loaded config: " + string(configBytes))
	log.Fatalln("Exiting Program on success...")
}