package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelRefundCloseAPIRequest) Reset() {
	r._closeRefundOrderRequest = nil
	r.Params.ToZero()
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

var poolAlibabaAscpChannelRefundCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelRefundCloseRequest()
	},
}

// GetAlibabaAscpChannelRefundCloseRequest 从 sync.Pool 获取 AlibabaAscpChannelRefundCloseAPIRequest
func GetAlibabaAscpChannelRefundCloseAPIRequest() *AlibabaAscpChannelRefundCloseAPIRequest {
	return poolAlibabaAscpChannelRefundCloseAPIRequest.Get().(*AlibabaAscpChannelRefundCloseAPIRequest)
}

// ReleaseAlibabaAscpChannelRefundCloseAPIRequest 将 AlibabaAscpChannelRefundCloseAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelRefundCloseAPIRequest(v *AlibabaAscpChannelRefundCloseAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelRefundCloseAPIRequest.Put(v)
}
