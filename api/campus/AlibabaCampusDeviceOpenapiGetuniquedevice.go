package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiGetuniquedevice 根据设备uuid获取设备信息
// alibaba.campus.device.openapi.getuniquedevice
//
// 根据设备uuid获取设备信息
func AlibabaCampusDeviceOpenapiGetuniquedevice(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiGetuniquedeviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
