package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigetdevicelistAPIRequest 多条件查询设备分组 API请求
// alibaba.campus.device.openapi.getdevicelist
//
// 多条件查询设备分组
type AlibabacampusdeviceopenapigetdevicelistAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 多条件查询对象
	_query *DeviceApiQuery
}

// NewAlibabacampusdeviceopenapigetdevicelistRequest 初始化AlibabacampusdeviceopenapigetdevicelistAPIRequest对象
func NewAlibabacampusdeviceopenapigetdevicelistRequest() *AlibabacampusdeviceopenapigetdevicelistAPIRequest {
	return &AlibabacampusdeviceopenapigetdevicelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigetdevicelistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getdevicelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigetdevicelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigetdevicelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabacampusdeviceopenapigetdevicelistAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapigetdevicelistAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 多条件查询对象
func (r *AlibabacampusdeviceopenapigetdevicelistAPIRequest) SetQuery(_query *DeviceApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabacampusdeviceopenapigetdevicelistAPIRequest) GetQuery() *DeviceApiQuery {
	return r._query
}
