package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkdeviceunifystatusset 设置设备标准属性状态
// alibaba.alink.device.unify.status.set
//
// 操作用户绑定的设备
func Alibabaalinkdeviceunifystatusset(clt *core.SDKClient, req *alink.AlibabaalinkdeviceunifystatussetAPIRequest, session string) (*alink.AlibabaalinkdeviceunifystatussetAPIResponse, error) {
	var resp alink.AlibabaalinkdeviceunifystatussetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
