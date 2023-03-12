package main

import (
	"flag"
	"log"
	"net"

	"github.com/lajunta/getip/grpcd"
	"google.golang.org/grpc"
)


var (
	port string
)

func parseFlags() {
	flag.StringVar(&port, "port", "5505", "运行端口")
	flag.Parse()
}

func init() {
	parseFlags()
	loadRooms()
}

func main() {

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	grpcd.RegisterGetIPServiceServer(s, &GetIPServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
