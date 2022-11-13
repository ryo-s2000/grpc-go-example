package main

import (
	"context"
	echo "grpc-go-example/pbs"
	"log"
)

// EchoService を実装するサーバーの構造体
type server struct {
	echo.UnimplementedEchoServiceServer
}

// EchoServiceServer インタフェースの Echo メソッドの実装（本物の Echo 実装）
func (s *server) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Printf("Received from client: %v", in.GetMessage())
	return &echo.EchoResponse{Message: "*" + in.GetMessage()}, nil
}
