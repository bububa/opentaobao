package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderRefundConfirmAPIRequest 退款确认 API请求
// alibaba.wdk.channel.order.refund.confirm
//
// 退款确认
type AlibabaWdkChannelOrderRefundConfirmAPIRequest struct {
	model.Params
	// 退款确认信息
	_orderRefundConfirmInfo *OrderRefundConfirmInfo
}

// NewAlibabaWdkChannelOrderRefundConfirmRequest 初始化AlibabaWdkChannelOrderRefundConfirmAPIRequest对象
func NewAlibabaWdkChannelOrderRefundConfirmRequest() *AlibabaWdkChannelOrderRefundConfirmAPIRequest {
	return &AlibabaWdkChannelOrderRefundConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkChannelOrderRefundConfirmAPIRequest) Reset() {
	r._orderRefundConfirmInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderRefundConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.refund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderRefundConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkChannelOrderRefundConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRefundConfirmInfo is OrderRefundConfirmInfo Setter
// 退款确认信息
func (r *AlibabaWdkChannelOrderRefundConfirmAPIRequest) SetOrderRefundConfirmInfo(_orderRefundConfirmInfo *OrderRefundConfirmInfo) error {
	r._orderRefundConfirmInfo = _orderRefundConfirmInfo
	r.Set("order_refund_confirm_info", _orderRefundConfirmInfo)
	return nil
}

// GetOrderRefundConfirmInfo OrderRefundConfirmInfo Getter
func (r AlibabaWdkChannelOrderRefundConfirmAPIRequest) GetOrderRefundConfirmInfo() *OrderRefundConfirmInfo {
	return r._orderRefundConfirmInfo
}

var poolAlibabaWdkChannelOrderRefundConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkChannelOrderRefundConfirmRequest()
	},
}

// GetAlibabaWdkChannelOrderRefundConfirmRequest 从 sync.Pool 获取 AlibabaWdkChannelOrderRefundConfirmAPIRequest
func GetAlibabaWdkChannelOrderRefundConfirmAPIRequest() *AlibabaWdkChannelOrderRefundConfirmAPIRequest {
	return poolAlibabaWdkChannelOrderRefundConfirmAPIRequest.Get().(*AlibabaWdkChannelOrderRefundConfirmAPIRequest)
}

// ReleaseAlibabaWdkChannelOrderRefundConfirmAPIRequest 将 AlibabaWdkChannelOrderRefundConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkChannelOrderRefundConfirmAPIRequest(v *AlibabaWdkChannelOrderRefundConfirmAPIRequest) {
	v.Reset()
	poolAlibabaWdkChannelOrderRefundConfirmAPIRequest.Put(v)
}
