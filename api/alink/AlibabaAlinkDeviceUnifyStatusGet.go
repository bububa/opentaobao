package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkdeviceunifystatusget 查询设备标准属性最新状态
// alibaba.alink.device.unify.status.get
//
// 查询设备最新标准属性状态
func Alibabaalinkdeviceunifystatusget(clt *core.SDKClient, req *alink.AlibabaalinkdeviceunifystatusgetAPIRequest, session string) (*alink.AlibabaalinkdeviceunifystatusgetAPIResponse, error) {
	var resp alink.AlibabaalinkdeviceunifystatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
