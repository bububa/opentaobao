package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest
saveeventinfoforibos API请求
alibaba.campus.device.openapi.saveeventinfoforibos

IBos的事件信息上报与反馈的处理接口 */
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 系统自动生成
	_param1 *EventInfoApiDto
}

// New
