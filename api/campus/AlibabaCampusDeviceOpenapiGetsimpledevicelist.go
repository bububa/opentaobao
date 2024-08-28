package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetsimpledevicelist 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
// alibaba.campus.device.openapi.getsimpledevicelist
//
// 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
func AlibabaCampusDeviceOpenapiGetsimpledevicelist(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
