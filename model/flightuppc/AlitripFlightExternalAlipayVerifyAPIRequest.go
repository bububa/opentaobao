package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightexternalalipayverifyAPIRequest 支付宝小程序验签 API请求
// alitrip.flight.external.alipay.verify
//
// 支付宝小程序验签
type AlitripflightexternalalipayverifyAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayCheckSignReq *AlipayCheckSignReq
}

// NewAlitripflightexternalalipayverifyRequest 初始化AlitripflightexternalalipayverifyAPIRequest对象
func NewAlitripflightexternalalipayverifyRequest() *AlitripflightexternalalipayverifyAPIRequest {
	return &AlitripflightexternalalipayverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightexternalalipayverifyAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightexternalalipayverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightexternalalipayverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayCheckSignReq is AlipayCheckSignReq Setter
// 入参结构体
func (r *AlitripflightexternalalipayverifyAPIRequest) SetAlipayCheckSignReq(_alipayCheckSignReq *AlipayCheckSignReq) error {
	r._alipayCheckSignReq = _alipayCheckSignReq
	r.Set("alipay_check_sign_req", _alipayCheckSignReq)
	return nil
}

// GetAlipayCheckSignReq AlipayCheckSignReq Getter
func (r AlitripflightexternalalipayverifyAPIRequest) GetAlipayCheckSignReq() *AlipayCheckSignReq {
	return r._alipayCheckSignReq
}
