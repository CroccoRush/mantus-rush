package simulation

import "github.com/golang/protobuf/ptypes/timestamp"

func SetTime(pbTime *timestamp.Timestamp) {
	Time = pbTime.Seconds
}

func GetPBTime() *timestamp.Timestamp {
	return &timestamp.Timestamp{Seconds: Time}
}
