package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkdeviceinfoupdate 更新设备昵称等信息
// alibaba.alink.device.info.update
//
// 更新设备昵称等信息
func Alibabaalinkdeviceinfoupdate(clt *core.SDKClient, req *alink.AlibabaalinkdeviceinfoupdateAPIRequest, session string) (*alink.AlibabaalinkdeviceinfoupdateAPIResponse, error) {
	var resp alink.AlibabaalinkdeviceinfoupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
