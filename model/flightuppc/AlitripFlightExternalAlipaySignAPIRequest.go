package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipaySignAPIRequest 支付宝小程序验签 API请求
// alitrip.flight.external.alipay.sign
//
// 支付宝小程序验签
type AlitripFlightExternalAlipaySignAPIRequest struct {
	model.Params
	// 入参结构体
	_alipaySignReq *AlipaySignReq
}

// NewAlitripFlightExternalAlipaySignRequest 初始化AlitripFlightExternalAlipaySignAPIRequest对象
func NewAlitripFlightExternalAlipaySignRequest() *AlitripFlightExternalAlipaySignAPIRequest {
	return &AlitripFlightExternalAlipaySignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightExternalAlipaySignAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightExternalAlipaySignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightExternalAlipaySignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipaySignReq is AlipaySignReq Setter
// 入参结构体
func (r *AlitripFlightExternalAlipaySignAPIRequest) SetAlipaySignReq(_alipaySignReq *AlipaySignReq) error {
	r._alipaySignReq = _alipaySignReq
	r.Set("alipay_sign_req", _alipaySignReq)
	return nil
}

// GetAlipaySignReq AlipaySignReq Getter
func (r AlitripFlightExternalAlipaySignAPIRequest) GetAlipaySignReq() *AlipaySignReq {
	return r._alipaySignReq
}
