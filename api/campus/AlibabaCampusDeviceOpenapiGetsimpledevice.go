package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetsimpledevice 获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
// alibaba.campus.device.openapi.getsimpledevice
//
// 获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
func AlibabaCampusDeviceOpenapiGetsimpledevice(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetsimpledeviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
