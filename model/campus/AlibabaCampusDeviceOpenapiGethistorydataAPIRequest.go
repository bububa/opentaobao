package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigethistorydataAPIRequest 查询设备历史数据 API请求
// alibaba.campus.device.openapi.gethistorydata
//
// 查询历史数据的接口
type AlibabacampusdeviceopenapigethistorydataAPIRequest struct {
	model.Params
	// 请求端信息
	_workBenchContext *WorkBenchContext
	// 历史数据查询对象
	_query *DeviceDataApiQuery
}

// NewAlibabacampusdeviceopenapigethistorydataRequest 初始化AlibabacampusdeviceopenapigethistorydataAPIRequest对象
func NewAlibabacampusdeviceopenapigethistorydataRequest() *AlibabacampusdeviceopenapigethistorydataAPIRequest {
	return &AlibabacampusdeviceopenapigethistorydataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigethistorydataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.gethistorydata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigethistorydataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigethistorydataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求端信息
func (r *AlibabacampusdeviceopenapigethistorydataAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapigethistorydataAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 历史数据查询对象
func (r *AlibabacampusdeviceopenapigethistorydataAPIRequest) SetQuery(_query *DeviceDataApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabacampusdeviceopenapigethistorydataAPIRequest) GetQuery() *DeviceDataApiQuery {
	return r._query
}
