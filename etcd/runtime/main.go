package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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
	var runtime_viper = viper.New()

	runtime_viper.AddRemoteProvider("etcd", "http://0.0.0.0:2379", "/config/config.yaml")
	runtime_viper.SetConfigType("yaml")
	if err := runtime_viper.ReadRemoteConfig(); err != nil {
		log.Fatal(err)
	}

	var c Config
	runtime_viper.Unmarshal(&c)

	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				fmt.Printf("unable to read remote config: %v\n", err)
				continue
			}

			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			runtime_viper.Unmarshal(&c)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, c)
	})

	http.ListenAndServe(":8080", nil)
}
