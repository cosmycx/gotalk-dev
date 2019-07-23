package main

import (
	"fmt"
	"io"
	"log"
	"net"

	greetpb "greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {

	fmt.Println("Greet Everyone function invoked with streaming request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "!"

		sErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})

		if sErr != nil {
			log.Fatalf("Error while sending: %v", sErr)
			return sErr
		}
	}
} // .GreetEveryone

// START OMIT
// func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {

func main() {
	fmt.Println("Hello, starting Greet RPC server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

} // .main
// END OMIT
