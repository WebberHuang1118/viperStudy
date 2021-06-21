package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	var runtime_viper = viper.New()

	runtime_viper.AddRemoteProvider("etcd", "http://0.0.0.0:2379", "/config/config.yaml")
	runtime_viper.SetConfigType("yaml")
	if err := runtime_viper.ReadRemoteConfig(); err != nil {
		log.Fatal(err)
	}

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
			// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// 	fmt.Fprint(w, runtime_viper.GetString("app_name"))
			// })
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, runtime_viper.GetString("app_name"))
	})

	http.ListenAndServe(":8080", nil)
}
