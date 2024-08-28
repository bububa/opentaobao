package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiSaveeventinfoforibos saveeventinfoforibos
// alibaba.campus.device.openapi.saveeventinfoforibos
//
// IBos的事件信息上报与反馈的处理接口
func AlibabaCampusDeviceOpenapiSaveeventinfoforibos(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
