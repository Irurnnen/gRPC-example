package controllers

import (
	"context"

	"github.com/Irurnnen/gRPCexample/internal/models"
	pb "github.com/Irurnnen/gRPCexample/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HelloController struct {
	pb.UnimplementedGreeterServer
}

var hm models.Hello

func (hc *HelloController) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// Get data from model
	reply, err := hm.SayHello(in.Name)

	// Process errors
	switch err {
	case nil:
		break
	case models.ErrExample:
		return &pb.HelloReply{}, status.Error(codes.OutOfRange, err.Error())
	default:
		// TODO: Create error example
		return &pb.HelloReply{}, nil
	}

	// Return data
	return &pb.HelloReply{Message: reply}, nil
}
