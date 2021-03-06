package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/abhishekratnam/say-grpc/api"

	"google.golang.org/grpc"
)

func main() {

	// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("could not listen to port %d: %v", *port, err)
	}
	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	s.Serve(lis)
	if err != nil {
		log.Fatalf("could not serve %v", err)
	}

}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	return nil, fmt.Errorf("not implemented")
}
