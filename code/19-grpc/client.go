package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	greetpb "greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("from client ...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doBiDiStreaming(c)
} // .main

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting a Client streaming RPC...")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jack",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jill",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	} // .requests

	// START OMIT
	// create stream
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	waitchan := make(chan struct{})

	// send to server | go routines
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second)
		} // .for
		stream.CloseSend()
	}() // .go.func
	// END OMIT

	// receive messages from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		} // .for
		close(waitchan)
	}()

	// wait for all done
	<-waitchan
} // .doBiDiStreaming
