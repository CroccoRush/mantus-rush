// Package pb2 implements a server for service.
package pb2

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"

	"user_server/api/pb2/pb"
	"user_server/internal/simulation"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUserServer
}

// GetShipsState method
func (s *server) GetShipsState(in *empty.Empty, out pb.User_GetShipsStateServer) error {
	for _, ship := range simulation.Ships {
		if err := out.Send(ship); err != nil {
			return err
		}
	}
	return nil
}

// GetShipState method
func (s *server) GetShipState(ctx context.Context, in *pb.ShipRequest) (*pb.ShipResponse, error) {
	ship := simulation.Ships[in.Name]
	fmt.Println(in.Name)
	return ship, nil
}

// SetShipState method
func (s *server) SetShipState(ctx context.Context, in *pb.SetShipRequest) (*pb.Response, error) {
	//ToDo
	fmt.Println(in.Name)
	ship := simulation.Ships[in.Name]
	ship.Velocity = in.Velocity
	for i, time := range ship.Way.Timestamp {
		if time == in.Timestamp {
			ship.Way.GeoPoint[i] = in.GeoPoint
			break
		}
	}
	return &pb.Response{Message: "point set", Timestamp: simulation.GetPBTime(), ErrorCode: 200}, nil
}

// SetShipSimWay method
func (s *server) SetShipSimWay(ctx context.Context, in *pb.SetShipSimRequest) (*pb.Response, error) {
	//ToDo
	fmt.Println(in.Name)

	return &pb.Response{Message: "Hello " + in.GetName(), Timestamp: simulation.GetPBTime()}, nil
}

// SetShipsSimWay method
func (s *server) SetShipsSimWay(out pb.User_SetShipsSimWayServer) error {
	//ToDo
	for {
		ship, err := out.Recv()
		if err != nil {
			log.Printf("error; %s", err)
		}
		if ship.Name == "-1" {
			break
		}
		simulation.Ships[ship.Name].Way = ship.Way
	}

	return nil
}

// NewShip method
func (s *server) NewShip(ctx context.Context, in *pb.NewShipRequest) (*pb.Response, error) {
	//ToDo
	log.Println(in.Name)
	in.
	return &pb.Response{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
