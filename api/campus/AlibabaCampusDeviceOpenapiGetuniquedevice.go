package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetuniquedevice 根据设备uuid获取设备信息
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
func AlibabaCampusDeviceOpenapiGetuniquedevice(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest, session string) (*campus.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse, error) {
	var resp campus.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
