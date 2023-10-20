package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightexternalalipaysignAPIRequest 支付宝小程序验签 API请求
// alitrip.flight.external.alipay.sign
//
// 支付宝小程序验签
type AlitripflightexternalalipaysignAPIRequest struct {
	model.Params
	// 入参结构体
	_alipaySignReq *AlipaySignReq
}

// NewAlitripflightexternalalipaysignRequest 初始化AlitripflightexternalalipaysignAPIRequest对象
func NewAlitripflightexternalalipaysignRequest() *AlitripflightexternalalipaysignAPIRequest {
	return &AlitripflightexternalalipaysignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightexternalalipaysignAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightexternalalipaysignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightexternalalipaysignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipaySignReq is AlipaySignReq Setter
// 入参结构体
func (r *AlitripflightexternalalipaysignAPIRequest) SetAlipaySignReq(_alipaySignReq *AlipaySignReq) error {
	r._alipaySignReq = _alipaySignReq
	r.Set("alipay_sign_req", _alipaySignReq)
	return nil
}

// GetAlipaySignReq AlipaySignReq Getter
func (r AlitripflightexternalalipaysignAPIRequest) GetAlipaySignReq() *AlipaySignReq {
	return r._alipaySignReq
}
