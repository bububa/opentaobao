package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatcwmsoutboundordercancelAPIRequest 取消出库单 API请求
// alibaba.tcwms.outbound.order.cancel
//
// 取消出库单
type AlibabatcwmsoutboundordercancelAPIRequest struct {
	model.Params
	// 参数
	_outboundCancelRequest *OutboundCancelRequest
}

// NewAlibabatcwmsoutboundordercancelRequest 初始化AlibabatcwmsoutboundordercancelAPIRequest对象
func NewAlibabatcwmsoutboundordercancelRequest() *AlibabatcwmsoutboundordercancelAPIRequest {
	return &AlibabatcwmsoutboundordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatcwmsoutboundordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tcwms.outbound.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatcwmsoutboundordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatcwmsoutboundordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutboundCancelRequest is OutboundCancelRequest Setter
// 参数
func (r *AlibabatcwmsoutboundordercancelAPIRequest) SetOutboundCancelRequest(_outboundCancelRequest *OutboundCancelRequest) error {
	r._outboundCancelRequest = _outboundCancelRequest
	r.Set("outbound_cancel_request", _outboundCancelRequest)
	return nil
}

// GetOutboundCancelRequest OutboundCancelRequest Getter
func (r AlibabatcwmsoutboundordercancelAPIRequest) GetOutboundCancelRequest() *OutboundCancelRequest {
	return r._outboundCancelRequest
}
