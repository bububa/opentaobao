package perfect

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundOrderCancelAPIRequest 取消出库单 API请求
// alibaba.tcwms.outbound.order.cancel
//
// 取消出库单
type AlibabaTcwmsOutboundOrderCancelAPIRequest struct {
	model.Params
	// 参数
	_outboundCancelRequest *OutboundCancelRequest
}

// NewAlibabaTcwmsOutboundOrderCancelRequest 初始化AlibabaTcwmsOutboundOrderCancelAPIRequest对象
func NewAlibabaTcwmsOutboundOrderCancelRequest() *AlibabaTcwmsOutboundOrderCancelAPIRequest {
	return &AlibabaTcwmsOutboundOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTcwmsOutboundOrderCancelAPIRequest) Reset() {
	r._outboundCancelRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTcwmsOutboundOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tcwms.outbound.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTcwmsOutboundOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTcwmsOutboundOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutboundCancelRequest is OutboundCancelRequest Setter
// 参数
func (r *AlibabaTcwmsOutboundOrderCancelAPIRequest) SetOutboundCancelRequest(_outboundCancelRequest *OutboundCancelRequest) error {
	r._outboundCancelRequest = _outboundCancelRequest
	r.Set("outbound_cancel_request", _outboundCancelRequest)
	return nil
}

// GetOutboundCancelRequest OutboundCancelRequest Getter
func (r AlibabaTcwmsOutboundOrderCancelAPIRequest) GetOutboundCancelRequest() *OutboundCancelRequest {
	return r._outboundCancelRequest
}

var poolAlibabaTcwmsOutboundOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTcwmsOutboundOrderCancelRequest()
	},
}

// GetAlibabaTcwmsOutboundOrderCancelRequest 从 sync.Pool 获取 AlibabaTcwmsOutboundOrderCancelAPIRequest
func GetAlibabaTcwmsOutboundOrderCancelAPIRequest() *AlibabaTcwmsOutboundOrderCancelAPIRequest {
	return poolAlibabaTcwmsOutboundOrderCancelAPIRequest.Get().(*AlibabaTcwmsOutboundOrderCancelAPIRequest)
}

// ReleaseAlibabaTcwmsOutboundOrderCancelAPIRequest 将 AlibabaTcwmsOutboundOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaTcwmsOutboundOrderCancelAPIRequest(v *AlibabaTcwmsOutboundOrderCancelAPIRequest) {
	v.Reset()
	poolAlibabaTcwmsOutboundOrderCancelAPIRequest.Put(v)
}
