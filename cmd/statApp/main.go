package main

import (
	"context"
	"flag"
	"fmt"
	api "github.com/statApp/pkg/api"
	statistic "github.com/statApp/pkg/statistic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	time2 "time"
)

var (
	port = flag.Int("port", 50050, "the server port")
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type server struct {
	api.UnimplementedStatAppServiceServer
}

func (s *server) DayPriceAvg(ctx context.Context, date *api.RequestDayPriceAvg) (*api.ResponseDayPriceAvg, error) {

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := api.NewFakeDataSourceAppServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time2.Second)
	defer cancel()

	r, err := c.GetData(ctx, &api.RequestCurrencyInfo{Name: "ethereum"})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	avg := statistic.PriceAvg(r, date.GetDate())
	return &api.ResponseDayPriceAvg{DayPriceAvg: avg}, nil
}

func startGrpcServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	api.RegisterStatAppServiceServer(s, &server{})
	reflection.Register(s)
	log.Print("start stat server")
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func main() {
	startGrpcServer()
}
