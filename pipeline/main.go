package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	rc := RoamingConfig{}
	viper.Unmarshal(&rc)
	fmt.Printf("rc:%v\n\n", rc)

	// b, err := json.Marshal(&rc)
	// if err != nil {
	// 	log.Fatal("Marshal failed: %v", err)
	// }
	// fmt.Printf("encoding byte:%v\n\n", string(b))

	// rcDecode := RoamingConfig{}
	// if err := json.Unmarshal(b, &rcDecode); err != nil {
	// 	log.Fatal("Unmarshal failed: %v", err)
	// }

	// fmt.Printf("rcDecode:%v\n", rcDecode)
}
