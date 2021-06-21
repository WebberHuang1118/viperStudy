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

func init() {
	// 绑定环境变量
	viper.BindEnv("log_level", "LOG_LEVEL") //$export LOG_LEVEL=INFO
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.Set("redis.port", 7381)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println("app_name: ", viper.GetString("app_name"))
	fmt.Println("log_level: ", viper.GetString("log_level"))

	fmt.Println("mysql database: ", viper.GetString("mysql.database"))
	fmt.Println("mysql ip: ", viper.GetString("mysql.ip"))
	fmt.Println("mysql password: ", viper.GetInt("mysql.password"))
	fmt.Println("mysql port: ", viper.GetInt("mysql.port"))
	fmt.Println("mysql user: ", viper.GetString("mysql.user"))

	fmt.Println("redis ip: ", viper.GetString("redis.ip"))
	fmt.Println("redis port: ", viper.GetInt("redis.port"))

	fmt.Println("server ports: ", viper.GetIntSlice("server.ports"))
	fmt.Println("server protocols: ", viper.GetStringSlice("server.protocols"))

	viper.WriteConfig()
}
