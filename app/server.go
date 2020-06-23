package app

import (
	"context"
	"github.com/soichisumi-sandbox/grpc-opencensus-sample/app/proto"
)

func NewServer() Server {
	m := make(map[string]string)
	return Server{
		data: m,
	}
}

type Server struct{
	data map[string]string
}

func (s Server) DeleteData(ctx context.Context, req *proto.DeleteDataRequest) (*proto.DeleteDataResponse, error) {
	delete(s.data, req.Key)
	return &proto.DeleteDataResponse{}, nil
}

func (s Server) SetData(ctx context.Context, req *proto.SetDataRequest) (*proto.SetDataResponse, error) {
	s.data[req.Key] = req.Value
	return &proto.SetDataResponse{}, nil
}

func (s Server) GetAllData(ctx context.Context, req *proto.GetAllDataRequest) (*proto.GetAllDataResponse, error) {
	res := make([]*proto.Pair, 0)
	for k, v := range s.data {
		res = append(res, &proto.Pair{
			Key: k,
			Value: v,
		})
	}
	return &proto.GetAllDataResponse{
		Result: res,
	}, nil
}

func (s Server) GetData(ctx context.Context, req *proto.GetDataRequest) (*proto.GetDataResponse, error) {
	return &proto.GetDataResponse{
		Value: s.data[req.Key],
	}, nil
}


