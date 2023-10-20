package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightexternalalipaydecryptAPIRequest 支付宝小程序密文解密 API请求
// alitrip.flight.external.alipay.decrypt
//
// 支付宝小程序密文解密
type AlitripflightexternalalipaydecryptAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayDecryptReq *AlipayDecryptReq
}

// NewAlitripflightexternalalipaydecryptRequest 初始化AlitripflightexternalalipaydecryptAPIRequest对象
func NewAlitripflightexternalalipaydecryptRequest() *AlitripflightexternalalipaydecryptAPIRequest {
	return &AlitripflightexternalalipaydecryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightexternalalipaydecryptAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightexternalalipaydecryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightexternalalipaydecryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayDecryptReq is AlipayDecryptReq Setter
// 入参结构体
func (r *AlitripflightexternalalipaydecryptAPIRequest) SetAlipayDecryptReq(_alipayDecryptReq *AlipayDecryptReq) error {
	r._alipayDecryptReq = _alipayDecryptReq
	r.Set("alipay_decrypt_req", _alipayDecryptReq)
	return nil
}

// GetAlipayDecryptReq AlipayDecryptReq Getter
func (r AlitripflightexternalalipaydecryptAPIRequest) GetAlipayDecryptReq() *AlipayDecryptReq {
	return r._alipayDecryptReq
}
