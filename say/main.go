package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/campoy/justforfunc/say-grpc/api"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file where the output written")
	flag.Parse()

	conn, err := net.Dial("tcp", *backend)
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()
	clinet := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text: "Hello"}
	res, err := clinet.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("Could not say %s: %v", err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("Could not write %s: %v", *output, err)
	}
}
