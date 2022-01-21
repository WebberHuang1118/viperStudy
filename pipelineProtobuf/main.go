package main

import (
	"log"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

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
		//log.Printf("bytes %v", string(bytes))

		return bytes, nil
	}
}

type Pool struct {
	Value int
	Sizes []int
}

type Private struct {
	Pools []Pool
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

	var rc RoamingConfig
	if err := viper.Unmarshal(&rc, optDecode); err != nil {
		log.Fatalf("viper unmarshal failed %v ", err)
	}
	log.Printf("rc:%v\n", rc)

	var p Private
	bs := rc.GetPipeline().GetServices()[1].GetPrivate()
	if err := yaml.Unmarshal(bs, &p); err != nil {
		log.Fatalf("Unmarshal private failed %v\n", err)
	}
	log.Printf("p:%v\n", p)
}
