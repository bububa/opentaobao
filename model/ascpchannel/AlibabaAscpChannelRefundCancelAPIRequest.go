package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelrefundcancelAPIRequest 渠道退款单撤销 API请求
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
type AlibabaascpchannelrefundcancelAPIRequest struct {
	model.Params
	// 入参
	_cancelRefundOrderRequest *Cancelrefundorderrequest
}

// NewAlibabaascpchannelrefundcancelRequest 初始化AlibabaascpchannelrefundcancelAPIRequest对象
func NewAlibabaascpchannelrefundcancelRequest() *AlibabaascpchannelrefundcancelAPIRequest {
	return &AlibabaascpchannelrefundcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelrefundcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.refund.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelrefundcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelrefundcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRefundOrderRequest is CancelRefundOrderRequest Setter
// 入参
func (r *AlibabaascpchannelrefundcancelAPIRequest) SetCancelRefundOrderRequest(_cancelRefundOrderRequest *Cancelrefundorderrequest) error {
	r._cancelRefundOrderRequest = _cancelRefundOrderRequest
	r.Set("cancel_refund_order_request", _cancelRefundOrderRequest)
	return nil
}

// GetCancelRefundOrderRequest CancelRefundOrderRequest Getter
func (r AlibabaascpchannelrefundcancelAPIRequest) GetCancelRefundOrderRequest() *Cancelrefundorderrequest {
	return r._cancelRefundOrderRequest
}
