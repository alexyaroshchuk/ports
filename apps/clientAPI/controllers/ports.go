package controllers

import "context"

type PortService interface {
	GetPort(context.Context, *pb.RequestWithId) (*pb.PortResponse, error)
}

func GetPort()  {
	ConvertPo
}