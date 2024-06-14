// Package pb2 implements a client for Greeter service.
package pb2

import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user_server/api/pb2/pb"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateState(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not update state: %v", err)
	}
	log.Printf("state updated: %s", r.GetMessage())
}
