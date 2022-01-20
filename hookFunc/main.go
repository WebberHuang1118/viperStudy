package main

import (
	"encoding/json"
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	//goflag "flag"
	"reflect"

	"github.com/c2h5oh/datasize"
	"github.com/mitchellh/mapstructure"
)

//var ip *int = flag.Int("flagname", 1234, "help message for flagname")

type Config struct {
	Server Server
}

type Data struct {
	Team string
	Name string
}

type Server struct {
	Id            string
	Tcp_Bind      string
	DashboardBind string
	MaxSize       *datasize.ByteSize
	Private       map[string][]byte
}

type transit struct {
	James   []byte
	Curry   []byte
	Country []byte
	Sizes   []byte
}

func main() {
	//var ip = flag.IntP(("flagname", "f", 1234, "help message")
	//flag.Lookup("flagname").NoOptDefVal = "4321"

	var ip = flag.String("flagname", "172.0.0.1", "help message for flagname")

	//goflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()
	fmt.Println("ip has value ", *ip)

	viper.SetConfigName("config") // 只有名字，沒有後綴
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 查看並重新讀取配置文件
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})
	viper.Set("server.Tcp_Bind", *ip)
	fmt.Printf("watch the tcp_bind : %s\n", viper.Get("server.TcpBind"))
	fmt.Printf("watch the dash_bind : %s\n", viper.Get("server.DashboardBind"))

	Cfg := Config{Server: Server{Private: make(map[string][]byte)}}
	//fmt.Println("watch the viper is : ", viper.GetViper())

	optDecode := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		StringToByteSizesHookFunc(),
		myHookFunc(),
	))

	err = viper.Unmarshal(&Cfg, optDecode)
	if err != nil {

	}

	//err = viper.Unmarshal(&Cfg)
	//if err != nil {
	//	fmt.Println("The error from Unmarshal is : ", err)
	//}
	fmt.Println("watch the config of Cfg private is : ", Cfg.Server.Private)
	//fmt.Printf("watch the config of MaxSize : %v\n", Cfg.Server.MaxSize)

	var t transit
	if err := mapstructure.WeakDecode(Cfg.Server.Private, &t); err != nil {
		log.Printf("Decode private failed %v\n", err)
	}
	//fmt.Println("transit Data is : ", t.Data)

	james := Data{}
	if err := json.Unmarshal([]byte(t.James), &james); err != nil {
		log.Printf("Unmarshal james failed %v\n", err)
	}
	log.Printf("james:%v\n", james)

	curry := Data{}
	if err := json.Unmarshal([]byte(t.Curry), &curry); err != nil {
		log.Printf("Unmarshal curry failed %v\n", err)
	}
	log.Printf("curry:%v\n", curry)

	var country string
	if err := json.Unmarshal([]byte(t.Country), &country); err != nil {
		log.Printf("Unmarshal country failed %v\n", err)
	}
	log.Printf("country:%v\n", country)

	sizes := []int{}
	if err := json.Unmarshal([]byte(t.Sizes), &sizes); err != nil {
		log.Printf("Unmarshal size failed %v\n", err)
	}
	log.Printf("size:%v\n", sizes)
}

func StringToByteSizesHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {

		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(datasize.ByteSize(5)) {
			return data, nil
		}

		// Convert it by parsing
		raw := data.(string)
		result := new(datasize.ByteSize)
		result.UnmarshalText([]byte(raw))
		return result.Bytes(), nil
	}
}

func myHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {

		// if f.Kind() != reflect.Map {
		// 	return data, nil
		// }

		if t != reflect.TypeOf(map[string][]byte{}) {
			return data, nil
		}

		fmt.Printf("f %v\n", f)
		fmt.Printf("t %v\n", t)

		// Convert it by parsing
		raw := data.(map[string]interface{})
		ret := map[string][]byte{}
		for k, v := range raw {
			fmt.Printf("k:%v v:%v\n", k, v)
			bytes, err := json.Marshal(v)
			if err != nil {
				fmt.Printf("marshal fail %v\n", err)
			}
			ret[k] = bytes
		}

		return ret, nil
	}
}
