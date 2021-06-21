package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type roamingConfig map[string]*ConnectConfig

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	var rc roamingConfig
	viper.Unmarshal(&rc)

	for k, v := range rc {
		fmt.Printf("\n---- stage:%v ----\n", k)
		fmt.Printf("initiator:%v\n", v.GetInitiator())
		if v.GetTrigger() != nil {
			fmt.Printf("trigger:%v\n", v.GetTrigger())
		}
	}
}
