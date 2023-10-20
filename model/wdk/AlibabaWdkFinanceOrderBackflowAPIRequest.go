package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfinanceorderbackflowAPIRequest 财务订单回流 API请求
// alibaba.wdk.finance.order.backflow
//
// 星巴克拉取财务订单回流数据
type AlibabawdkfinanceorderbackflowAPIRequest struct {
	model.Params
	// 财务订单回流入参
	_financeOrderDetailRequest *FinanceOrderDetailRequest
}

// NewAlibabawdkfinanceorderbackflowRequest 初始化AlibabawdkfinanceorderbackflowAPIRequest对象
func NewAlibabawdkfinanceorderbackflowRequest() *AlibabawdkfinanceorderbackflowAPIRequest {
	return &AlibabawdkfinanceorderbackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfinanceorderbackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.finance.order.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfinanceorderbackflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfinanceorderbackflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFinanceOrderDetailRequest is FinanceOrderDetailRequest Setter
// 财务订单回流入参
func (r *AlibabawdkfinanceorderbackflowAPIRequest) SetFinanceOrderDetailRequest(_financeOrderDetailRequest *FinanceOrderDetailRequest) error {
	r._financeOrderDetailRequest = _financeOrderDetailRequest
	r.Set("finance_order_detail_request", _financeOrderDetailRequest)
	return nil
}

// GetFinanceOrderDetailRequest FinanceOrderDetailRequest Getter
func (r AlibabawdkfinanceorderbackflowAPIRequest) GetFinanceOrderDetailRequest() *FinanceOrderDetailRequest {
	return r._financeOrderDetailRequest
}
