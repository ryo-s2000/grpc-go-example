package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	echo "grpc-go-example/pbs"
)

const addr = "localhost"
const port = "52000"
const url = addr + ":" + port

func main() {
	// EchoService サーバーへ接続する
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := echo.NewEchoServiceClient(conn)

	// Echo メソッドを呼び出す
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &echo.EchoRequest{Message: "AAAAA"})
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}
	log.Printf("Received from server: %s", r.GetMessage())
}
