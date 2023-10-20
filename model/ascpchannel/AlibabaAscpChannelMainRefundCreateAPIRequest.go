package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelMainRefundCreateAPIRequest 淘外分销逆向创单（未发货整单退） API请求
// alibaba.ascp.channel.main.refund.create
//
// 淘外分销解决方案--订单--逆向创单（未发货整单退）
type AlibabaAscpChannelMainRefundCreateAPIRequest struct {
	model.Params
	// 逆向单创建请求
	_refundCreateRequest *ExternalCreateRefundOrderRequest
}

// NewAlibabaAscpChannelMainRefundCreateRequest 初始化AlibabaAscpChannelMainRefundCreateAPIRequest对象
func NewAlibabaAscpChannelMainRefundCreateRequest() *AlibabaAscpChannelMainRefundCreateAPIRequest {
	return &AlibabaAscpChannelMainRefundCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelMainRefundCreateAPIRequest) Reset() {
	r._refundCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelMainRefundCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.main.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelMainRefundCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelMainRefundCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCreateRequest is RefundCreateRequest Setter
// 逆向单创建请求
func (r *AlibabaAscpChannelMainRefundCreateAPIRequest) SetRefundCreateRequest(_refundCreateRequest *ExternalCreateRefundOrderRequest) error {
	r._refundCreateRequest = _refundCreateRequest
	r.Set("refund_create_request", _refundCreateRequest)
	return nil
}

// GetRefundCreateRequest RefundCreateRequest Getter
func (r AlibabaAscpChannelMainRefundCreateAPIRequest) GetRefundCreateRequest() *ExternalCreateRefundOrderRequest {
	return r._refundCreateRequest
}

var poolAlibabaAscpChannelMainRefundCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelMainRefundCreateRequest()
	},
}

// GetAlibabaAscpChannelMainRefundCreateRequest 从 sync.Pool 获取 AlibabaAscpChannelMainRefundCreateAPIRequest
func GetAlibabaAscpChannelMainRefundCreateAPIRequest() *AlibabaAscpChannelMainRefundCreateAPIRequest {
	return poolAlibabaAscpChannelMainRefundCreateAPIRequest.Get().(*AlibabaAscpChannelMainRefundCreateAPIRequest)
}

// ReleaseAlibabaAscpChannelMainRefundCreateAPIRequest 将 AlibabaAscpChannelMainRefundCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelMainRefundCreateAPIRequest(v *AlibabaAscpChannelMainRefundCreateAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelMainRefundCreateAPIRequest.Put(v)
}
