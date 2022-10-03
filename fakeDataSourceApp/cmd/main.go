package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	api "github.com/statApp/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 50051, "the server port")
)

const (
	dataPath = "./fakeDataSourceApp/gas_price.json"
)

type server struct {
	api.UnimplementedFakeDataSourceAppServiceServer
}

func (s *server) GetData(ctx context.Context, name *api.RequestCurrencyInfo) (*api.ResponseCurrencyInfo, error) {
	if name.Name == "ethereum" {
		file, err := os.Open(dataPath)
		defer file.Close()

		if err != nil {
			return nil, err
		}

		dataBytes, err := ioutil.ReadAll(file)

		var data api.ResponseCurrencyInfo

		err = json.Unmarshal(dataBytes, &data)

		if err != nil {
			return nil, err
		}

		return &data, nil
	}
	return nil, errors.New("incorrect currency name")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	api.RegisterFakeDataSourceAppServiceServer(s, &server{})
	reflection.Register(s)
	log.Print("start fake server")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
