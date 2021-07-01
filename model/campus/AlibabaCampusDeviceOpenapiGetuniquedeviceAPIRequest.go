package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest
根据设备uuid获取设备信息 API请求
alibaba.campus.device.openapi.getuniquedevice

根据设备uuid获取设备信息 */
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备序列号uuid
	_uuid string
}

// New
