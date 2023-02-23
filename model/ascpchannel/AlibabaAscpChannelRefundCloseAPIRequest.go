package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelRefundCloseAPIRequest 渠道退款单关闭 API请求
// alibaba.ascp.channel.refund.close
//
// 渠道退款单关闭
type AlibabaAscpChannelRefundCloseAPIRequest struct {
	model.Params
	// 入参
	_closeRefundOrderRequest *Closerefundorderrequest
}

// NewAlibabaAscpChannelRefundCloseRequest 初始化AlibabaAscpChannelRefundCloseAPIRequest对象
func NewAlibabaAscpChannelRefundCloseRequest() *AlibabaAscpChannelRefundCloseAPIRequest {
	return &AlibabaAscpChannelRefundCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelRefundCloseAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.refund.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelRefundCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelRefundCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloseRefundOrderRequest is CloseRefundOrderRequest Setter
// 入参
func (r *AlibabaAscpChannelRefundCloseAPIRequest) SetCloseRefundOrderRequest(_closeRefundOrderRequest *Closerefundorderrequest) error {
	r._closeRefundOrderRequest = _closeRefundOrderRequest
	r.Set("close_refund_order_request", _closeRefundOrderRequest)
	return nil
}

// GetCloseRefundOrderRequest CloseRefundOrderRequest Getter
func (r AlibabaAscpChannelRefundCloseAPIRequest) GetCloseRefundOrderRequest() *Closerefundorderrequest {
	return r._closeRefundOrderRequest
}
