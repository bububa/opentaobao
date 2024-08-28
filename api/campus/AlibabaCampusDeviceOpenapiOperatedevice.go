package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiOperatedevice 根据uuid操作设备
// alibaba.campus.device.openapi.operatedevice
//
// 根据uuid操作设备
func AlibabaCampusDeviceOpenapiOperatedevice(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiOperatedeviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
