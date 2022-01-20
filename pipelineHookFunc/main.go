package main

import (
	"log"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Pipeline Pipeline
}

type Pipeline struct {
	Services []Service
}

type Service struct {
	Name     string
	Uri      string
	Private  []byte
	Actioner Actioner
}

type Actioner struct {
	Uri string
}

type Pool struct {
	Value int
	Sizes []int
}

type Private struct {
	Pools []Pool
}

func myHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {

		if t != reflect.TypeOf([]byte{}) {
			return data, nil
		}

		// fmt.Printf("f %v\n", f)
		// fmt.Printf("t %v\n", t)
		// fmt.Printf("data %v\n", data)

		bytes, err := yaml.Marshal(data)
		if err != nil {
			log.Fatalf("marshal data fail %v", err)
		}
		log.Printf("bytes %v", string(bytes))

		return bytes, nil
	}
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	optDecode := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(myHookFunc()))

	cfg := Config{}
	if err := viper.Unmarshal(&cfg, optDecode); err != nil {
		log.Fatalf("viper unmarshal failed %v ", err)
	}
	log.Printf("cfg:%v\n", cfg)
	log.Printf("private:%v\n", string(cfg.Pipeline.Services[1].Private))

	var p Private
	if err := yaml.Unmarshal(cfg.Pipeline.Services[1].Private, &p); err != nil {
		log.Fatalf("Unmarshal private failed %v\n", err)
	}
	log.Printf("p:%v\n", p)
}
