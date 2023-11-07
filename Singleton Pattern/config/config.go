package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	Environment string `json:"environment"`
	Port string `json:"port"`
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host string `json:"host"`
		Port string `json:"port"`
		Schema string `json:"schema"`
	}
	Logger string `json:"logger"`
}

var instance *config
var once sync.Once
func GetConfig() *config {
	once.Do(func ()  {
		instance = &config{}
		instance.InitConfig()
	})
	return instance
}

var mu sync.Mutex
func (c *config) InitConfig() {
	mu.Lock()
	defer mu.Unlock()
	
	viper.AddConfigPath(".")
    viper.SetConfigName("config")
    viper.SetConfigType("json")

    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
		log.Println("Error: " + err.Error())
		return
    }

    err = viper.Unmarshal(&c)
	if err != nil {
		log.Println("Error: " + err.Error())
		return
	}

	log.Printf("Successfully loaded config: %v\n", c)
    return
}