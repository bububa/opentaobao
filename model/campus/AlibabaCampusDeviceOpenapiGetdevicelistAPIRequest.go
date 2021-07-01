package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest
多条件查询设备分组 API请求
alibaba.campus.device.openapi.getdevicelist

多条件查询设备分组 */
type AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 多条件查询对象
	_query *DeviceApiQuery
}

// New
