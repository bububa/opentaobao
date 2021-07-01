package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest
根据设备uuid获取设备采集信息 API请求
alibaba.campus.device.openapi.getdevicerealtimelog

根据设备uuid获取设备采集信息 */
type AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备uuid
	_uuid string
}

// New
