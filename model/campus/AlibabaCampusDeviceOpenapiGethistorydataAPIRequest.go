package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGethistorydataAPIRequest
查询设备历史数据 API请求
alibaba.campus.device.openapi.gethistorydata

查询历史数据的接口 */
type AlibabaCampusDeviceOpenapiGethistorydataAPIRequest struct {
	model.Params
	// 请求端信息
	_workBenchContext *WorkBenchContext
	// 历史数据查询对象
	_query *DeviceDataApiQuery
}

// New
