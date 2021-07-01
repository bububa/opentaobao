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

// NewAlibabaCampusDeviceOpenapiGetsimpledevicelistRequest 初始化AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetsimpledevicelistRequest() *AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getsimpledevicelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// Get WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// Set is Query Setter
// 多条件查询对象
func (r *AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest) SetQuery(_query *DeviceApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r AlibabaCampusDeviceOpenapiGetsimpledevicelistAPIRequest) GetQuery() *DeviceApiQuery {
	return r._query
}
