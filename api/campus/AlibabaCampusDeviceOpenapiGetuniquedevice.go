package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetuniquedevice 根据设备uuid获取设备信息
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
func AlibabaCampusDeviceOpenapiGetuniquedevice(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
