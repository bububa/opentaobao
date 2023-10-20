package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkdevicedetailget 获取设备详情
// alibaba.alink.device.detail.get
//
// 阿里智能获取设备详情
func Alibabaalinkdevicedetailget(clt *core.SDKClient, req *alink.AlibabaalinkdevicedetailgetAPIRequest, session string) (*alink.AlibabaalinkdevicedetailgetAPIResponse, error) {
	var resp alink.AlibabaalinkdevicedetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
