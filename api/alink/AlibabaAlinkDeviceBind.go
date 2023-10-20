package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkdevicebind 绑定设备
// alibaba.alink.device.bind
//
// 阿里智能解绑设备
func Alibabaalinkdevicebind(clt *core.SDKClient, req *alink.AlibabaalinkdevicebindAPIRequest, session string) (*alink.AlibabaalinkdevicebindAPIResponse, error) {
	var resp alink.AlibabaalinkdevicebindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
