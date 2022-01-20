package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName  string
	LogLevel string

	MySQLS []MySQLConfig
	Redis  RedisConfig
	Server ServerConfig
}

type MySQLConfig struct {
	Database string
	IP       string
	Password string
	Port     int
	User     string
}

type Locations struct {
	Locations []MySQLConfig
}

type RedisConfig struct {
	IP   string
	Port int
}

type ServerConfig struct {
	Ports     []int
	Protocols []string
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	//mysqls := []MySQLConfig{}
	subCfg := viper.Sub("mysqls")

	c := subCfg.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal sub config to YAML: %v", err)
	}
	fmt.Printf("bs:\n%v\n", string(bs))

	// viper.UnmarshalKey("mysqls", &mysqls)
	// fmt.Println(mysqls)

	loc := Locations{Locations: []MySQLConfig{}}
	// if err := subCfg.Unmarshal(&loc); err != nil {
	// 	log.Fatalf("unable to unmarshal sub config to struct: %v", err)
	// }
	if err != yaml.Unmarshal(bs, &loc) {
		log.Fatalf("unable to unmarshal binary to loc: %v", err)
	}
	fmt.Println(loc)
}
