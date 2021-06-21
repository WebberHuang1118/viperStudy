package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string `mapstructure:"app_name"`
	LogLevel string `mapstructure:"log_level"`

	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Server ServerConfig `mapstructure:"server"`
}

type MySQLConfig struct {
	Database string `mapstructure:"database"`
	IP       string `mapstructure:"ip"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
}

type RedisConfig struct {
	IP   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Ports     []int    `mapstructure:"ports"`
	Protocols []string `mapstructure:"protocols"`
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.Set("redis.port", 5381)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	var c Config
	viper.Unmarshal(&c)

	fmt.Println(c)
}
