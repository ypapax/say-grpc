package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	pb "github.com/ypapax/say-grpc/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	logrus.Infof("listening to port %+v", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}
	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	if err = s.Serve(lis); err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}

}

type server struct{}

func (s server) Say(c context.Context, t *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create a temp file")
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close temporary file")
	}

	cmd := exec.Command("flite", "-t", t.Text, "-o", f.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("unable to run flite command: %+v", err)
	}
	b, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("unable to read output file: %+v", err)
	}
	return &pb.Speech{Audio: b}, nil
}
