package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceHistorydataGetAPIRequest 设备历史数据批量获取 API请求
// alibaba.campus.device.historydata.get
//
// 设备历史数据批量获取
type AlibabaCampusDeviceHistorydataGetAPIRequest struct {
	model.Params
	// workbench
	_workBenchContext *WorkBenchContext
	// 查询条件
	_query *DeviceHistoryBatchQuery
}

// NewAlibabaCampusDeviceHistorydataGetRequest 初始化AlibabaCampusDeviceHistorydataGetAPIRequest对象
func NewAlibabaCampusDeviceHistorydataGetRequest() *AlibabaCampusDeviceHistorydataGetAPIRequest {
	return &AlibabaCampusDeviceHistorydataGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceHistorydataGetAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.historydata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceHistorydataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceHistorydataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// workbench
func (r *AlibabaCampusDeviceHistorydataGetAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceHistorydataGetAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 查询条件
func (r *AlibabaCampusDeviceHistorydataGetAPIRequest) SetQuery(_query *DeviceHistoryBatchQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaCampusDeviceHistorydataGetAPIRequest) GetQuery() *DeviceHistoryBatchQuery {
	return r._query
}
