package flightuppc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipaySnQueryAPIRequest 支付宝小程序查询证书序列号 API请求
// alitrip.flight.external.alipay.sn.query
//
// 支付宝小程序查询证书序列号
type AlitripFlightExternalAlipaySnQueryAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayQueryCertSnReq *AlipayQueryCertSnReq
}

// NewAlitripFlightExternalAlipaySnQueryRequest 初始化AlitripFlightExternalAlipaySnQueryAPIRequest对象
func NewAlitripFlightExternalAlipaySnQueryRequest() *AlitripFlightExternalAlipaySnQueryAPIRequest {
	return &AlitripFlightExternalAlipaySnQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightExternalAlipaySnQueryAPIRequest) Reset() {
	r._alipayQueryCertSnReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightExternalAlipaySnQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.sn.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightExternalAlipaySnQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightExternalAlipaySnQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayQueryCertSnReq is AlipayQueryCertSnReq Setter
// 入参结构体
func (r *AlitripFlightExternalAlipaySnQueryAPIRequest) SetAlipayQueryCertSnReq(_alipayQueryCertSnReq *AlipayQueryCertSnReq) error {
	r._alipayQueryCertSnReq = _alipayQueryCertSnReq
	r.Set("alipay_query_cert_sn_req", _alipayQueryCertSnReq)
	return nil
}

// GetAlipayQueryCertSnReq AlipayQueryCertSnReq Getter
func (r AlitripFlightExternalAlipaySnQueryAPIRequest) GetAlipayQueryCertSnReq() *AlipayQueryCertSnReq {
	return r._alipayQueryCertSnReq
}

var poolAlitripFlightExternalAlipaySnQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightExternalAlipaySnQueryRequest()
	},
}

// GetAlitripFlightExternalAlipaySnQueryRequest 从 sync.Pool 获取 AlitripFlightExternalAlipaySnQueryAPIRequest
func GetAlitripFlightExternalAlipaySnQueryAPIRequest() *AlitripFlightExternalAlipaySnQueryAPIRequest {
	return poolAlitripFlightExternalAlipaySnQueryAPIRequest.Get().(*AlitripFlightExternalAlipaySnQueryAPIRequest)
}

// ReleaseAlitripFlightExternalAlipaySnQueryAPIRequest 将 AlitripFlightExternalAlipaySnQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightExternalAlipaySnQueryAPIRequest(v *AlitripFlightExternalAlipaySnQueryAPIRequest) {
	v.Reset()
	poolAlitripFlightExternalAlipaySnQueryAPIRequest.Put(v)
}
