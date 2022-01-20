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

	MySQL  MySQLConfig
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

	//viper.Set("redis.port", 5381)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	// var c Config
	// viper.Unmarshal(&c)

	// fmt.Println(c)

	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	fmt.Printf("bs:%v\n", string(bs))

	var cfg Config
	if err != yaml.Unmarshal(bs, &cfg) {
		log.Fatalf("unable to unmarshal binary to config: %v", err)
	}
	fmt.Printf("cfg:%v\n", cfg)
}
