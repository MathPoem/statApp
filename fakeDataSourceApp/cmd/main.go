package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pkg "github.com/statApp/pkg/api"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const (
	dataPath = "github.com/MathPoem/statApp/fakeDataSourceApp/gas_price.json"
	port     = ":50500"
)

type server struct {
	pkg.UnimplementedFakeDataSourceAppServiceServer
}

func (s *server) GetData(ctx context.Context, name *pkg.RequestCurrencyInfo) (*pkg.ResponseCurrencyInfo, error) {
	if name.Name == "ethereum" {
		file, err := os.Open(dataPath)
		defer file.Close()

		if err != nil {
			return nil, err
		}

		dataBytes, err := ioutil.ReadAll(file)

		var data pkg.ResponseCurrencyInfo

		err = json.Unmarshal(dataBytes, &data)

		if err != nil {
			return nil, err
		}

		return &data, nil
	}
	return nil, errors.New("incorrect currency name")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pkg.RegisterFakeDataSourceAppServiceServer(s, &server{})
	fmt.Println("start fake server")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
