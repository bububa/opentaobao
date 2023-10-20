package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelrefundcloseAPIRequest 渠道退款单关闭 API请求
// alibaba.ascp.channel.refund.close
//
// 渠道退款单关闭
type AlibabaascpchannelrefundcloseAPIRequest struct {
	model.Params
	// 入参
	_closeRefundOrderRequest *Closerefundorderrequest
}

// NewAlibabaascpchannelrefundcloseRequest 初始化AlibabaascpchannelrefundcloseAPIRequest对象
func NewAlibabaascpchannelrefundcloseRequest() *AlibabaascpchannelrefundcloseAPIRequest {
	return &AlibabaascpchannelrefundcloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelrefundcloseAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.refund.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelrefundcloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelrefundcloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloseRefundOrderRequest is CloseRefundOrderRequest Setter
// 入参
func (r *AlibabaascpchannelrefundcloseAPIRequest) SetCloseRefundOrderRequest(_closeRefundOrderRequest *Closerefundorderrequest) error {
	r._closeRefundOrderRequest = _closeRefundOrderRequest
	r.Set("close_refund_order_request", _closeRefundOrderRequest)
	return nil
}

// GetCloseRefundOrderRequest CloseRefundOrderRequest Getter
func (r AlibabaascpchannelrefundcloseAPIRequest) GetCloseRefundOrderRequest() *Closerefundorderrequest {
	return r._closeRefundOrderRequest
}
