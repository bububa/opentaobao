package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdevicehistorydatagetAPIRequest 设备历史数据批量获取 API请求
// alibaba.campus.device.historydata.get
//
// 设备历史数据批量获取
type AlibabacampusdevicehistorydatagetAPIRequest struct {
	model.Params
	// workbench
	_workBenchContext *WorkBenchContext
	// 查询条件
	_query *DeviceHistoryBatchQuery
}

// NewAlibabacampusdevicehistorydatagetRequest 初始化AlibabacampusdevicehistorydatagetAPIRequest对象
func NewAlibabacampusdevicehistorydatagetRequest() *AlibabacampusdevicehistorydatagetAPIRequest {
	return &AlibabacampusdevicehistorydatagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdevicehistorydatagetAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.historydata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdevicehistorydatagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdevicehistorydatagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// workbench
func (r *AlibabacampusdevicehistorydatagetAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdevicehistorydatagetAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 查询条件
func (r *AlibabacampusdevicehistorydatagetAPIRequest) SetQuery(_query *DeviceHistoryBatchQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabacampusdevicehistorydatagetAPIRequest) GetQuery() *DeviceHistoryBatchQuery {
	return r._query
}
