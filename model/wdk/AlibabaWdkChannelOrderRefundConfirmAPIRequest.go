package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelorderrefundconfirmAPIRequest 退款确认 API请求
// alibaba.wdk.channel.order.refund.confirm
//
// 退款确认
type AlibabawdkchannelorderrefundconfirmAPIRequest struct {
	model.Params
	// 退款确认信息
	_orderRefundConfirmInfo *OrderRefundConfirmInfo
}

// NewAlibabawdkchannelorderrefundconfirmRequest 初始化AlibabawdkchannelorderrefundconfirmAPIRequest对象
func NewAlibabawdkchannelorderrefundconfirmRequest() *AlibabawdkchannelorderrefundconfirmAPIRequest {
	return &AlibabawdkchannelorderrefundconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkchannelorderrefundconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.refund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkchannelorderrefundconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkchannelorderrefundconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRefundConfirmInfo is OrderRefundConfirmInfo Setter
// 退款确认信息
func (r *AlibabawdkchannelorderrefundconfirmAPIRequest) SetOrderRefundConfirmInfo(_orderRefundConfirmInfo *OrderRefundConfirmInfo) error {
	r._orderRefundConfirmInfo = _orderRefundConfirmInfo
	r.Set("order_refund_confirm_info", _orderRefundConfirmInfo)
	return nil
}

// GetOrderRefundConfirmInfo OrderRefundConfirmInfo Getter
func (r AlibabawdkchannelorderrefundconfirmAPIRequest) GetOrderRefundConfirmInfo() *OrderRefundConfirmInfo {
	return r._orderRefundConfirmInfo
}
