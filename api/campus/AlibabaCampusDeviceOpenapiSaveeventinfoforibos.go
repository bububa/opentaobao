package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiSaveeventinfoforibos saveeventinfoforibos
// alibaba.campus.device.openapi.saveeventinfoforibos
//
// IBos的事件信息上报与反馈的处理接口
func AlibabaCampusDeviceOpenapiSaveeventinfoforibos(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
