package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest
查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API请求
alibaba.campus.device.openapi.getsimpledevicelist

查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) */
type AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 多条件查询对象
	_query *DeviceApiQuery
}

// New
