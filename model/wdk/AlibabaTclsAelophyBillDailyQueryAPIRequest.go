package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophybilldailyqueryAPIRequest 账单日汇总接口 API请求
// alibaba.tcls.aelophy.bill.daily.query
//
// 账单日汇总接口
type AlibabatclsaelophybilldailyqueryAPIRequest struct {
	model.Params
	// 请求入参
	_dailyRequest *BillDailyQueryRequest
}

// NewAlibabatclsaelophybilldailyqueryRequest 初始化AlibabatclsaelophybilldailyqueryAPIRequest对象
func NewAlibabatclsaelophybilldailyqueryRequest() *AlibabatclsaelophybilldailyqueryAPIRequest {
	return &AlibabatclsaelophybilldailyqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophybilldailyqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.bill.daily.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophybilldailyqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophybilldailyqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDailyRequest is DailyRequest Setter
// 请求入参
func (r *AlibabatclsaelophybilldailyqueryAPIRequest) SetDailyRequest(_dailyRequest *BillDailyQueryRequest) error {
	r._dailyRequest = _dailyRequest
	r.Set("daily_request", _dailyRequest)
	return nil
}

// GetDailyRequest DailyRequest Getter
func (r AlibabatclsaelophybilldailyqueryAPIRequest) GetDailyRequest() *BillDailyQueryRequest {
	return r._dailyRequest
}
