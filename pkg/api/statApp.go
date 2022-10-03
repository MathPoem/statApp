package pkg

import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	time2 "time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type Server struct {
	UnimplementedStatAppServiceServer
}

func (s *Server) GetStat(ctx context.Context, empty *empty.Empty) (*ResponseStat, error) {

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := NewFakeDataSourceAppServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time2.Second)
	defer cancel()

	r, err := c.GetData(ctx, empty)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stat := Stat(r)
	return stat, nil
}
