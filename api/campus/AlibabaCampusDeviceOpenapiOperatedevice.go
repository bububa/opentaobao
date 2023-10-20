package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiOperatedevice 根据uuid操作设备
// alibaba.campus.device.openapi.operatedevice
//
// 根据uuid操作设备
func AlibabaCampusDeviceOpenapiOperatedevice(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
