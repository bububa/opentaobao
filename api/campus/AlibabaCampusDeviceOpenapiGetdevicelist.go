package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetdevicelist 多条件查询设备分组
// alibaba.campus.device.openapi.getdevicelist
//
// 多条件查询设备分组
func AlibabaCampusDeviceOpenapiGetdevicelist(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
