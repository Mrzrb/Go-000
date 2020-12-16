package internal

import (
	"context"
	"main/api"
	"main/internal/biz"
)

type DemoService struct {
	B *biz.DemoBiz
}

func NewDemoService(b *biz.DemoBiz) *DemoService {
	return &DemoService{B: b}
}

func (d *DemoService) Demo(ctx context.Context, req *api.GetDemoReq) (*api.GetDemoRsp, error) {
	datas := []*api.Demo{{Id: 123, Blob: "sad"}}
	rsp := &api.GetDemoRsp{
		Data: datas,
	}
	return rsp, nil
}
