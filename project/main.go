package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/mfsyahrz/submission/project/db"
	"github.com/mfsyahrz/submission/project/internal/config"
	"github.com/mfsyahrz/submission/project/internal/infrastructure/omdb"
	postgres "github.com/mfsyahrz/submission/project/internal/infrastructure/postgres/search_activity"
	GRPC "github.com/mfsyahrz/submission/project/internal/transport/grpc"
	"github.com/mfsyahrz/submission/project/internal/transport/rest"
	"github.com/mfsyahrz/submission/project/internal/usecase/movie"
	"google.golang.org/grpc"
)

func main() {

	cfg, err := config.New(".env")
	if err != nil {
		panic("Failed to setup config: " + err.Error())
	}

	httpAddr := flag.String("http", cfg.Service.RestAddress(), "http listen address")

	errChan := make(chan error)

	e := echo.New()

	db, err := db.Initialize(cfg.Postgres)
	if err != nil {
		panic("Failed to init database: " + err.Error())
	}
	activityRepo := postgres.New(db.Conn)
	movieDBintg := omdb.New(&cfg.OMDB)
	u := movie.New(activityRepo, movieDBintg)
	rest.NewMovieHandler(e, u)

	grpcSrv := grpc.NewServer()
	grpcHandler := GRPC.NewHandler(u)

	listen, err := net.Listen("tcp", cfg.Service.GrpcAddress())
	if err != nil {
		panic("Failed to Listen: " + err.Error())
	}

	GRPC.RegisterMovieServiceServer(grpcSrv, grpcHandler.MovieHandler)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go startRestServer(e, httpAddr)

	go startGRPCServer(cfg, grpcSrv, listen)

	err = <-errChan
	panic("failed to start server" + err.Error())

}

func startRestServer(e *echo.Echo, httpAddr *string) {
	if err := e.Start(*httpAddr); err != nil {
		panic("Failed to serve REST: " + err.Error())
	}
}

func startGRPCServer(cfg *config.Config, grpcSrv *grpc.Server, listen net.Listener) {
	fmt.Println("grpc listen on :" + cfg.Service.Port.GRPC)
	if err := grpcSrv.Serve(listen); err != nil {
		panic("Failed to serve GRPC: " + err.Error())
	}
}
