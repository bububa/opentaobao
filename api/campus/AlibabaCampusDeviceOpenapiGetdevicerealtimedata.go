package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetdevicerealtimedata 获取指定设备下指定参数的实时值
// alibaba.campus.device.openapi.getdevicerealtimedata
//
// 获取指定设备下指定参数的实时值
func AlibabaCampusDeviceOpenapiGetdevicerealtimedata(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
