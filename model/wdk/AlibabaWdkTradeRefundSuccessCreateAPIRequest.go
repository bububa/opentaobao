package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktraderefundsuccesscreateAPIRequest 五道口终态逆向订单创建 API请求
// alibaba.wdk.trade.refund.success.create
//
// 五道口终态逆向订单创建
type AlibabawdktraderefundsuccesscreateAPIRequest struct {
	model.Params
	// 逆向单请求对象
	_refundOrderRequest *AfterRefundOrderRequest
}

// NewAlibabawdktraderefundsuccesscreateRequest 初始化AlibabawdktraderefundsuccesscreateAPIRequest对象
func NewAlibabawdktraderefundsuccesscreateRequest() *AlibabawdktraderefundsuccesscreateAPIRequest {
	return &AlibabawdktraderefundsuccesscreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktraderefundsuccesscreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.success.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktraderefundsuccesscreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktraderefundsuccesscreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundOrderRequest is RefundOrderRequest Setter
// 逆向单请求对象
func (r *AlibabawdktraderefundsuccesscreateAPIRequest) SetRefundOrderRequest(_refundOrderRequest *AfterRefundOrderRequest) error {
	r._refundOrderRequest = _refundOrderRequest
	r.Set("refund_order_request", _refundOrderRequest)
	return nil
}

// GetRefundOrderRequest RefundOrderRequest Getter
func (r AlibabawdktraderefundsuccesscreateAPIRequest) GetRefundOrderRequest() *AfterRefundOrderRequest {
	return r._refundOrderRequest
}
