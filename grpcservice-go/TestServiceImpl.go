package helloworld

import (
	"context"
)

type TestServiceImpl struct {

}

func (*TestServiceImpl) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message:req.Name + "xxxxxxxx"}, nil
}
func (*TestServiceImpl) SayRepeatHello(req *RepeatHelloRequest, srv Greeter_SayRepeatHelloServer) error {
	for i:=0;i< int(req.Count);i++{
		srv.Send(&HelloReply{Message:req.Name + "xxx" + string(i)})
	}
	return nil
}
func (*TestServiceImpl) SayHelloAfterDelay(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message:req.Name + "xxxxxxxxDelay"}, nil
}
