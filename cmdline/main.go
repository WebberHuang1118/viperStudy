package main

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
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

func init() {
	pflag.Int("redis.port", 8381, "Redis port to connect")

	// 绑定命令行
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	pflag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	var c Config
	viper.Unmarshal(&c)

	fmt.Println(c)
}
