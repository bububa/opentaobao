package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipayDecryptAPIRequest 支付宝小程序密文解密 API请求
// alitrip.flight.external.alipay.decrypt
//
// 支付宝小程序密文解密
type AlitripFlightExternalAlipayDecryptAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayDecryptReq *AlipayDecryptReq
}

// NewAlitripFlightExternalAlipayDecryptRequest 初始化AlitripFlightExternalAlipayDecryptAPIRequest对象
func NewAlitripFlightExternalAlipayDecryptRequest() *AlitripFlightExternalAlipayDecryptAPIRequest {
	return &AlitripFlightExternalAlipayDecryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightExternalAlipayDecryptAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightExternalAlipayDecryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightExternalAlipayDecryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayDecryptReq is AlipayDecryptReq Setter
// 入参结构体
func (r *AlitripFlightExternalAlipayDecryptAPIRequest) SetAlipayDecryptReq(_alipayDecryptReq *AlipayDecryptReq) error {
	r._alipayDecryptReq = _alipayDecryptReq
	r.Set("alipay_decrypt_req", _alipayDecryptReq)
	return nil
}

// GetAlipayDecryptReq AlipayDecryptReq Getter
func (r AlitripFlightExternalAlipayDecryptAPIRequest) GetAlipayDecryptReq() *AlipayDecryptReq {
	return r._alipayDecryptReq
}
