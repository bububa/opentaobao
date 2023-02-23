package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyBillDailyQueryAPIRequest 账单日汇总接口 API请求
// alibaba.tcls.aelophy.bill.daily.query
//
// 账单日汇总接口
type AlibabaTclsAelophyBillDailyQueryAPIRequest struct {
	model.Params
	// 请求入参
	_dailyRequest *BillDailyQueryRequest
}

// NewAlibabaTclsAelophyBillDailyQueryRequest 初始化AlibabaTclsAelophyBillDailyQueryAPIRequest对象
func NewAlibabaTclsAelophyBillDailyQueryRequest() *AlibabaTclsAelophyBillDailyQueryAPIRequest {
	return &AlibabaTclsAelophyBillDailyQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyBillDailyQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.bill.daily.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyBillDailyQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyBillDailyQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDailyRequest is DailyRequest Setter
// 请求入参
func (r *AlibabaTclsAelophyBillDailyQueryAPIRequest) SetDailyRequest(_dailyRequest *BillDailyQueryRequest) error {
	r._dailyRequest = _dailyRequest
	r.Set("daily_request", _dailyRequest)
	return nil
}

// GetDailyRequest DailyRequest Getter
func (r AlibabaTclsAelophyBillDailyQueryAPIRequest) GetDailyRequest() *BillDailyQueryRequest {
	return r._dailyRequest
}
