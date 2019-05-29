package main

import (
	"fmt"
	"io"
	"log"
	api "recognize/go"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewDetectorClient(conn)
	stream, err := c.Detect(context.Background(), &api.DetectorRequest{})

	waitc := make(chan struct{})

	//Keep reading ....
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("EOF")
			}
			log.Printf("X : %d , Y : %d , W : %d , H : %d \n", in.X, in.Y, in.W, in.H)
		}
	}()
	<-waitc
	stream.CloseSend()

}
