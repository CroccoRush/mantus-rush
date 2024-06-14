package simulation

import "user_server/api/pb2/pb"

var (
	Time  int64
	Ships map[string]*pb.ShipResponse
)

func init() {
	Time = 0
	Ships = make(map[string]*pb.ShipResponse)
}
