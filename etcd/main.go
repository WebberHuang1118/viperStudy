package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	viper.AddRemoteProvider("etcd", "http://0.0.0.0:2379", "/config/config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("app_name: ", viper.GetString("app_name"))
}
