package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelRefundCancelAPIRequest 渠道退款单撤销 API请求
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
type AlibabaAscpChannelRefundCancelAPIRequest struct {
	model.Params
	// 入参
	_cancelRefundOrderRequest *Cancelrefundorderrequest
}

// NewAlibabaAscpChannelRefundCancelRequest 初始化AlibabaAscpChannelRefundCancelAPIRequest对象
func NewAlibabaAscpChannelRefundCancelRequest() *AlibabaAscpChannelRefundCancelAPIRequest {
	return &AlibabaAscpChannelRefundCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelRefundCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.refund.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelRefundCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCancelRefundOrderRequest is CancelRefundOrderRequest Setter
// 入参
func (r *AlibabaAscpChannelRefundCancelAPIRequest) SetCancelRefundOrderRequest(_cancelRefundOrderRequest *Cancelrefundorderrequest) error {
	r._cancelRefundOrderRequest = _cancelRefundOrderRequest
	r.Set("cancel_refund_order_request", _cancelRefundOrderRequest)
	return nil
}

// GetCancelRefundOrderRequest CancelRefundOrderRequest Getter
func (r AlibabaAscpChannelRefundCancelAPIRequest) GetCancelRefundOrderRequest() *Cancelrefundorderrequest {
	return r._cancelRefundOrderRequest
}
