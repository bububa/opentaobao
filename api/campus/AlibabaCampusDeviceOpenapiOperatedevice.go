package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapioperatedevice 根据uuid操作设备
// alibaba.campus.device.openapi.operatedevice
//
// 根据uuid操作设备
func Alibabacampusdeviceopenapioperatedevice(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapioperatedeviceAPIRequest, session string) (*campus.AlibabacampusdeviceopenapioperatedeviceAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapioperatedeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
