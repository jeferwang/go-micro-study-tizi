package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"jeferwang.com/tizi/proto"
	"time"
)

var service micro.Service

type OrderService struct {
}

func (o *OrderService) GetOrder(ctx context.Context, request *proto.GetOrderRequest, response *proto.Order) error {
	userId := 54

	userSrv := proto.NewUserSrvService("go.micro.api.userservice", service.Client())

	user, err := userSrv.GetAccount(context.TODO(), &proto.GetAccountRequest{Id: int32(userId)})

	if err == nil {
		response.Username = user.Username
		response.Address = user.Address
		response.Phone = user.Phone
	}

	response.Id = request.Id
	response.Name = "可乐"
	response.Price = 4.0
	response.CreateTime = time.Now().Format(time.RFC3339)
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

	// 定义一个微服务
	service = micro.NewService(micro.Name("go.micro.api.orderservice"), micro.Registry(reg))

	service.Init()

	proto.RegisterOrderSrvHandler(service.Server(), new(OrderService))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
