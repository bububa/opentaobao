package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetdevicerealtimelog 根据设备uuid获取设备采集信息
// alibaba.campus.device.openapi.getdevicerealtimelog
//
// 根据设备uuid获取设备采集信息
func AlibabaCampusDeviceOpenapiGetdevicerealtimelog(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
