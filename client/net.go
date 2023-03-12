package main

import (
	"flag"
	// "github.com/lajunta/getip/grpcd"
)

var (
	serverIP string
	port     string
)

func initArgs() {
	flag.StringVar(&serverIP, "server", "127.0.0.1", "server ip")
	flag.StringVar(&port, "port", "5505", "server port")
	flag.Parse()
}

/*

func net() {
	initArgs()
	conn, err := grpc.Dial(serverIP+":"+port, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("连接不成功 : %v", err)
	}

	defer conn.Close()

	c := grpcd.NewGetIPServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := grpcd.IPRequest{
		Room: "303",
		Mac:  "00:23:33:00:21",
	}

	res, err := c.GetIP(ctx, &req)

	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}

	fmt.Printf("%v", res)
}
*/
