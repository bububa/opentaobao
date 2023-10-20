package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息) API请求
// alibaba.campus.device.openapi.getsimpledevicelist
//
// 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
type AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 多条件查询对象
	_query *DeviceApiQuery
}

// NewAlibabacampusdeviceopenapigetsimpledevicelistRequest 初始化AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest对象
func NewAlibabacampusdeviceopenapigetsimpledevicelistRequest() *AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest {
	return &AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getsimpledevicelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 多条件查询对象
func (r *AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest) SetQuery(_query *DeviceApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest) GetQuery() *DeviceApiQuery {
	return r._query
}
