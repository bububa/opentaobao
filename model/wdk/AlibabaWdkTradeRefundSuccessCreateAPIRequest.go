package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundSuccessCreateAPIRequest 五道口终态逆向订单创建 API请求
// alibaba.wdk.trade.refund.success.create
//
// 五道口终态逆向订单创建
type AlibabaWdkTradeRefundSuccessCreateAPIRequest struct {
	model.Params
	// 逆向单请求对象
	_refundOrderRequest *AfterRefundOrderRequest
}

// NewAlibabaWdkTradeRefundSuccessCreateRequest 初始化AlibabaWdkTradeRefundSuccessCreateAPIRequest对象
func NewAlibabaWdkTradeRefundSuccessCreateRequest() *AlibabaWdkTradeRefundSuccessCreateAPIRequest {
	return &AlibabaWdkTradeRefundSuccessCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundSuccessCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.success.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundSuccessCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeRefundSuccessCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundOrderRequest is RefundOrderRequest Setter
// 逆向单请求对象
func (r *AlibabaWdkTradeRefundSuccessCreateAPIRequest) SetRefundOrderRequest(_refundOrderRequest *AfterRefundOrderRequest) error {
	r._refundOrderRequest = _refundOrderRequest
	r.Set("refund_order_request", _refundOrderRequest)
	return nil
}

// GetRefundOrderRequest RefundOrderRequest Getter
func (r AlibabaWdkTradeRefundSuccessCreateAPIRequest) GetRefundOrderRequest() *AfterRefundOrderRequest {
	return r._refundOrderRequest
}
