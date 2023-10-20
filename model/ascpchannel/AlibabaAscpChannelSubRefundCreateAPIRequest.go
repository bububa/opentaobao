package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSubRefundCreateAPIRequest 淘外分销逆向创单（子单退） API请求
// alibaba.ascp.channel.sub.refund.create
//
// 淘外分销逆向创单（子单退）
type AlibabaAscpChannelSubRefundCreateAPIRequest struct {
	model.Params
	// 子单退款创建请求
	_subRefundCreateReq *ExternalCreateRefundOrderDetailRequest
}

// NewAlibabaAscpChannelSubRefundCreateRequest 初始化AlibabaAscpChannelSubRefundCreateAPIRequest对象
func NewAlibabaAscpChannelSubRefundCreateRequest() *AlibabaAscpChannelSubRefundCreateAPIRequest {
	return &AlibabaAscpChannelSubRefundCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelSubRefundCreateAPIRequest) Reset() {
	r._subRefundCreateReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSubRefundCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sub.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSubRefundCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelSubRefundCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubRefundCreateReq is SubRefundCreateReq Setter
// 子单退款创建请求
func (r *AlibabaAscpChannelSubRefundCreateAPIRequest) SetSubRefundCreateReq(_subRefundCreateReq *ExternalCreateRefundOrderDetailRequest) error {
	r._subRefundCreateReq = _subRefundCreateReq
	r.Set("sub_refund_create_req", _subRefundCreateReq)
	return nil
}

// GetSubRefundCreateReq SubRefundCreateReq Getter
func (r AlibabaAscpChannelSubRefundCreateAPIRequest) GetSubRefundCreateReq() *ExternalCreateRefundOrderDetailRequest {
	return r._subRefundCreateReq
}

var poolAlibabaAscpChannelSubRefundCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelSubRefundCreateRequest()
	},
}

// GetAlibabaAscpChannelSubRefundCreateRequest 从 sync.Pool 获取 AlibabaAscpChannelSubRefundCreateAPIRequest
func GetAlibabaAscpChannelSubRefundCreateAPIRequest() *AlibabaAscpChannelSubRefundCreateAPIRequest {
	return poolAlibabaAscpChannelSubRefundCreateAPIRequest.Get().(*AlibabaAscpChannelSubRefundCreateAPIRequest)
}

// ReleaseAlibabaAscpChannelSubRefundCreateAPIRequest 将 AlibabaAscpChannelSubRefundCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelSubRefundCreateAPIRequest(v *AlibabaAscpChannelSubRefundCreateAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelSubRefundCreateAPIRequest.Put(v)
}
