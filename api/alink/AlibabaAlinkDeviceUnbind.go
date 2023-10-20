package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkdeviceunbind 解绑设备
// alibaba.alink.device.unbind
//
// 阿里智能解绑设备
func Alibabaalinkdeviceunbind(clt *core.SDKClient, req *alink.AlibabaalinkdeviceunbindAPIRequest, session string) (*alink.AlibabaalinkdeviceunbindAPIResponse, error) {
	var resp alink.AlibabaalinkdeviceunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
