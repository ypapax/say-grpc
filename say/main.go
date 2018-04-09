package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"

	"context"
	"fmt"
	"io/ioutil"
	"os"

	pb "github.com/ypapax/say-grpc/api"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav output file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	text := flag.Arg(0)

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to the server: %s: %+v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	res, err := client.Say(context.Background(), &pb.Text{Text: text})
	if err != nil {
		log.Fatalf("could not say %s: %+v", text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not write to file %s: %+v", *output, err)
	}

}
