package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest
获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API请求
alibaba.campus.device.openapi.getsimpledevice

获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) */
type AlibabaCampusDeviceOpenapiGetsimpledeviceAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备uuid
	_uuid string
}

// New
