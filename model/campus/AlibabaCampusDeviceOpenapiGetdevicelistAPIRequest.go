package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest 多条件查询设备分组 API请求
// alibaba.campus.device.openapi.getdevicelist
//
// 多条件查询设备分组
type AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 多条件查询对象
	_query *DeviceApiQuery
}

// NewAlibabaCampusDeviceOpenapiGetdevicelistRequest 初始化AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetdevicelistRequest() *AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getdevicelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 多条件查询对象
func (r *AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest) SetQuery(_query *DeviceApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaCampusDeviceOpenapiGetdevicelistAPIRequest) GetQuery() *DeviceApiQuery {
	return r._query
}
