package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"

	pb "github.com/ypapax/say-grpc/api"
	"context"
	"io/ioutil"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	text := flag.String("t", "hello", "string to say")
	output := flag.String("o", "output.wav", "wav output file")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to the server: %s: %+v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	res, err := client.Say(context.Background(), &pb.Text{Text: *text})
	if err != nil {
		log.Fatalf("could not say %s: %+v", *text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not write to file %s: %+v", *output, err)
	}

}