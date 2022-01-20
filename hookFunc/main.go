package main

import (
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	//goflag "flag"
	"reflect"

	"github.com/c2h5oh/datasize"
	"github.com/mitchellh/mapstructure"
)

//var ip *int = flag.Int("flagname", 1234, "help message for flagname")

type Config struct {
	Server Server
}

type Player struct {
	Name   string
	Number int
	Team   []string
}

type NBA struct {
	Players []Player
	Country string
	Sizes   []int
}

type Server struct {
	Id            string
	Tcp_Bind      string
	DashboardBind string
	MaxSize       *datasize.ByteSize
	Private       []byte
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
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	// 查看並重新讀取配置文件
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})
	viper.Set("server.Tcp_Bind", *ip)
	fmt.Printf("watch the tcp_bind : %s\n", viper.Get("server.TcpBind"))
	fmt.Printf("watch the dash_bind : %s\n", viper.Get("server.DashboardBind"))

	Cfg := Config{}
	//fmt.Println("watch the viper is : ", viper.GetViper())

	optDecode := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		StringToByteSizesHookFunc(),
		myHookFunc(),
	))

	err = viper.Unmarshal(&Cfg, optDecode)
	if err != nil {
		fmt.Println("The error from Unmarshal is : ", err)
	}

	//fmt.Println("watch the config of Cfg private is : ", Cfg.Server.Private)

	nba := NBA{Players: []Player{}}
	if err := yaml.Unmarshal(Cfg.Server.Private, &nba); err != nil {
		log.Fatalf("Unmarshal james failed %v\n", err)
	}
	log.Printf("nba:%v\n", nba)

	// nba := NBA{}
	// if err := mapstructure.WeakDecode(Cfg.Server.Private, &nba); err != nil {
	// 	log.Printf("Decode private failed %v\n", err)
	// }
	// log.Printf("nba:%v\n", nba)

	// var t transit
	// if err := mapstructure.WeakDecode(Cfg.Server.Private, &t); err != nil {
	// 	log.Printf("Decode private failed %v\n", err)
	// }
	// //fmt.Println("transit Data is : ", t.Data)

	// james := Data{}
	// if err := json.Unmarshal([]byte(t.James), &james); err != nil {
	// 	log.Printf("Unmarshal james failed %v\n", err)
	// }
	// log.Printf("james:%v\n", james)

	// curry := Data{}
	// if err := json.Unmarshal([]byte(t.Curry), &curry); err != nil {
	// 	log.Printf("Unmarshal curry failed %v\n", err)
	// }
	// log.Printf("curry:%v\n", curry)

	// var country string
	// if err := json.Unmarshal([]byte(t.Country), &country); err != nil {
	// 	log.Printf("Unmarshal country failed %v\n", err)
	// }
	// log.Printf("country:%v\n", country)

	// sizes := []int{}
	// if err := json.Unmarshal([]byte(t.Sizes), &sizes); err != nil {
	// 	log.Printf("Unmarshal size failed %v\n", err)
	// }
	// log.Printf("size:%v\n", sizes)
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

		return bytes, nil
	}
}
