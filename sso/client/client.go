package client

import (
	proto "github.com/Futturi/AuthSer/protos"
	"google.golang.org/grpc"
)

func NewGRPCClient(addr string) (proto.AuthClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewAuthClient(conn)
	return c, nil
}
