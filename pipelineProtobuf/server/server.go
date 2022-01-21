package main

import (
	"context"
	"log"
	"net"

	"github.com/webber/viperStudy/pipelineProtobuf/data"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

const (
	port = ":50051"
)

type Team struct {
	Name string
	City string
}

type League struct {
	Name     string
	Teams    []Team
	Country  string
	Division []string
}

type Country struct {
	Name       string
	Population int
	Capital    string
	Language   []string
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	data.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *data.RoamingConfig) (*data.RoamingConfig, error) {
	var league League
	bs := in.GetPipeline().GetServices()[1].GetPrivate()
	if err := yaml.Unmarshal(bs, &league); err != nil {
		log.Fatalf("Unmarshal private failed %v\n", err)
	}
	log.Printf("Received league:%v\n", league)

	var country Country
	bs = in.GetPipeline().GetServices()[1].GetActioner().GetPrivate()
	if err := yaml.Unmarshal(bs, &country); err != nil {
		log.Fatalf("Unmarshal private failed %v\n", err)
	}
	log.Printf("Received country:%v\n", country)
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	data.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
