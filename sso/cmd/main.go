package main

import (
	"log/slog"
	"net"

	proto "github.com/Futturi/AuthSer/protos"
	"github.com/Futturi/AuthSer/sso/internal/repository"
	"github.com/Futturi/AuthSer/sso/internal/server"
	"github.com/Futturi/AuthSer/sso/internal/service"
	"github.com/Futturi/AuthSer/sso/pkg"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	err := InitConfig()
	if err != nil {
		slog.Info(err.Error())
	}
	cfg := pkg.Config{
		Hostname: viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DB:       viper.GetString("db.namedb"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	db, err := pkg.InitPostges(cfg)
	if err != nil {
		slog.Info(err.Error())
	}
	repo := repository.NewRepository(db)
	servi := service.NewService(repo)
	server := server.NewServer(servi)

	ln, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		slog.Info(err.Error())
	}

	opts := []grpc.ServerOption{}
	serv := grpc.NewServer(opts...)
	proto.RegisterAuthServer(serv, server)

	if err != nil {
		slog.Info(err.Error())
	}
	if err := serv.Serve(ln); err != nil {
		slog.Info(err.Error())
	}
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}
