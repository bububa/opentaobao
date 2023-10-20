package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigetuniquedevice 根据设备uuid获取设备信息
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
func Alibabacampusdeviceopenapigetuniquedevice(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigetuniquedeviceAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigetuniquedeviceAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigetuniquedeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
