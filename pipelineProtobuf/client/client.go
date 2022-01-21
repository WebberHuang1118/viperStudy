/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/webber/viperStudy/pipelineProtobuf/data"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
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

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	optDecode := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(myHookFunc()))
	var rc data.RoamingConfig
	if err := viper.Unmarshal(&rc, optDecode); err != nil {
		log.Fatalf("viper unmarshal failed %v ", err)
	}
	//log.Printf("rc:%v\n", rc)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := data.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &rc)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
