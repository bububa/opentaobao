package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightexternalalipaysnqueryAPIRequest 支付宝小程序查询证书序列号 API请求
// alitrip.flight.external.alipay.sn.query
//
// 支付宝小程序查询证书序列号
type AlitripflightexternalalipaysnqueryAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayQueryCertSnReq *AlipayQueryCertSnReq
}

// NewAlitripflightexternalalipaysnqueryRequest 初始化AlitripflightexternalalipaysnqueryAPIRequest对象
func NewAlitripflightexternalalipaysnqueryRequest() *AlitripflightexternalalipaysnqueryAPIRequest {
	return &AlitripflightexternalalipaysnqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightexternalalipaysnqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.sn.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightexternalalipaysnqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightexternalalipaysnqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayQueryCertSnReq is AlipayQueryCertSnReq Setter
// 入参结构体
func (r *AlitripflightexternalalipaysnqueryAPIRequest) SetAlipayQueryCertSnReq(_alipayQueryCertSnReq *AlipayQueryCertSnReq) error {
	r._alipayQueryCertSnReq = _alipayQueryCertSnReq
	r.Set("alipay_query_cert_sn_req", _alipayQueryCertSnReq)
	return nil
}

// GetAlipayQueryCertSnReq AlipayQueryCertSnReq Getter
func (r AlitripflightexternalalipaysnqueryAPIRequest) GetAlipayQueryCertSnReq() *AlipayQueryCertSnReq {
	return r._alipayQueryCertSnReq
}
