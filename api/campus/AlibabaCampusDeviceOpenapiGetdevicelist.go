package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetdevicelist 多条件查询设备分组
// alibaba.campus.device.openapi.getdevicelist
//
// 多条件查询设备分组
func AlibabaCampusDeviceOpenapiGetdevicelist(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetdevicelistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
