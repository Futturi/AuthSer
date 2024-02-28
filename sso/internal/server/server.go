package server

import (
	"context"

	proto "github.com/Futturi/AuthSer/protos"
	"github.com/Futturi/AuthSer/sso/internal/service"
)

type Server struct {
	service *service.Service
	proto.UnimplementedAuthServer
}

func NewServer(service *service.Service) *Server {
	return &Server{service: service}
}

func (p *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	id, err := p.service.Register(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.RegisterResponse{UserId: int64(id)}, nil
}

func (p *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	token, err := p.service.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{Token: token}, nil
}
