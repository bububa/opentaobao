package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdevicegetdeviceforqueryAPIRequest 下发设备的分页接口(无需AOP控制) API请求
// alibaba.campus.device.getdeviceforquery
//
// 下发设备的分页接口(发布在TOP上，connect调用，无需AOP控制)
type AlibabacampusdevicegetdeviceforqueryAPIRequest struct {
	model.Params
	// 平台统一参数
	_workBenchContext *WorkBenchContext
	// 系统自动生成
	_query *DeviceApiQuery
}

// NewAlibabacampusdevicegetdeviceforqueryRequest 初始化AlibabacampusdevicegetdeviceforqueryAPIRequest对象
func NewAlibabacampusdevicegetdeviceforqueryRequest() *AlibabacampusdevicegetdeviceforqueryAPIRequest {
	return &AlibabacampusdevicegetdeviceforqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdevicegetdeviceforqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.getdeviceforquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdevicegetdeviceforqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdevicegetdeviceforqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 平台统一参数
func (r *AlibabacampusdevicegetdeviceforqueryAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdevicegetdeviceforqueryAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetQuery is Query Setter
// 系统自动生成
func (r *AlibabacampusdevicegetdeviceforqueryAPIRequest) SetQuery(_query *DeviceApiQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabacampusdevicegetdeviceforqueryAPIRequest) GetQuery() *DeviceApiQuery {
	return r._query
}
