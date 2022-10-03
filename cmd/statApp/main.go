package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	fakeService "github.com/statApp/fakeDataSourceApp/cmd"
	pkg "github.com/statApp/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"path"
	"strings"
)

var (
	grpcPort        = flag.Int("grpc-port", 50050, "port for gRPC server to listen on")
	httpPort        = flag.Int("http-port", 50040, "port for HTTP server to listen on")
	swaggerSpecPath = flag.String("swagger-spec", "./api/openapiv2/api/statApp.swagger.json", "path to Swagger spec file")
	staticFilesPath = flag.String("static-files", "./static/web", "path to static files")
)

func registerGatewayEndpoints() (http.Handler, error) {
	h := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pkg.RegisterStatAppServiceHandlerFromEndpoint(context.Background(), h, fmt.Sprintf(":%d", *grpcPort), opts); err != nil {
		return nil, fmt.Errorf("couldn't register HTTP handler: %w", err)
	}
	return h, nil
}

func startHTTPServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *swaggerSpecPath)
	})
	mux.HandleFunc("/swagger-ui/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(*staticFilesPath, "index.html"))
	})
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, "/static")
		p = path.Join(*staticFilesPath, p)
		http.ServeFile(w, r, p)
	})

	gw, err := registerGatewayEndpoints()
	if err != nil {
		log.Fatalf("couldn't register gateway endpoints: %v", err)
	}

	mux.Handle("/", gw)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), mux); err != nil {
		log.Fatalf("HTTP server stopped with error: %v", err)
	}
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pkg.RegisterStatAppServiceServer(s, &pkg.Server{})
	reflection.Register(s)
	log.Print("start stat server")
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	go fakeService.StartFakeDataSource()
	go startHTTPServer()
	startGrpcServer()
}
