package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"jeferwang.com/tizi/proto"
	"github.com/micro/go-micro/registry"
)

type UserService struct {
}

func (u *UserService) GetAccount(ctx context.Context, request *proto.GetAccountRequest, response *proto.Account) error {
	response.Id = request.Id
	response.Username = "jeferwang"
	response.Address = "ShenZhen"
	response.Phone = "18839136081"
	return nil
}

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:32379",
			"127.0.0.1:32381",
			"127.0.0.1:32383",
		}
	})

	service := micro.NewService(micro.Name("go.micro.api.userservice"), micro.Registry(reg))
	service.Init()

	proto.RegisterUserSrvHandler(service.Server(), new(UserService))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
