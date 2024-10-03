//go:build release
// +build release

package main

import (
	"fmt"
	"net"

	"github.com/Irurnnen/gRPCexample/internal/config"
	"github.com/Irurnnen/gRPCexample/internal/controllers"
	pb "github.com/Irurnnen/gRPCexample/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Setup logger
	zapConfig, err := config.GetZapConfig()
	if err != nil {
		fmt.Println("Error while getting zap log config" + err.Error())
		return
	}
	logging, err := zapConfig.Build()
	if err != nil {
		fmt.Println("Error while configuring zap logging" + err.Error())
		return
	}
	zap.ReplaceGlobals(logging)
	zap.S().Info("Zap started without errors")

	// TODO: SETUP SERVER

	lis, err := net.Listen("tcp", config.GetServer())
	if err != nil {
		zap.S().Errorw("failed listen", "host:port", config.GetServer(), "err", err)
	}
	// Create server
	s := grpc.NewServer()

	// Setup controllers
	helloController := new(controllers.HelloController)

	// Connect all controllers
	pb.RegisterGreeterServer(s, helloController)

	zap.S().Info("Server started without errors")
	zap.S().Debugw("Path co server", "host:port", config.GetServer())

	// Start server
	err = s.Serve(lis)
	if err != nil {
		zap.S().Errorw("failed to serve", "host:port", config.GetServer(), "err", err)
	}

}
