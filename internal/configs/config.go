package configs

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	parsedConfig string
	pathToConfig string = "./config.toml"
)

func init() {
	flag.StringVar(&parsedConfig, "config-file", pathToConfig, "Parses sqlite3 DSN and MQTT topic")
}

//Config keeps parsed sqlite3 dns and mqtt topic
type Config struct {
	DNS   string `toml:"dns"`
	Topic string `toml:"topic"`
}

//MakeConfig parses config file
func NewConfig() *Config{
	flag.Parse()

	config := new(Config)

	//try to decode config
	_, err := toml.DecodeFile(parsedConfig, &config)
	if err != nil {
		log.Println("Error while decoding config. New config with basic parameters will be created.")
		if err := createConfigFile(); err != nil {
			log.Fatal(err)
		}

		_, err := toml.DecodeFile(parsedConfig, &config)
		if err != nil {
			log.Fatal(err)
		}
	}
	return config
}

//createConfigFile creates config file with basic parameters
//dns = "users.db"
//topic = "test/topic"
func createConfigFile() error {
	settings := []byte("dns=\"users.db\"\ntopic=\"test/topic\"")

	err := ioutil.WriteFile("./config.toml", settings, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
