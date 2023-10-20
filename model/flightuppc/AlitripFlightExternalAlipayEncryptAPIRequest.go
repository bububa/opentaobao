package flightuppc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipayEncryptAPIRequest 支付宝小程序明文加密 API请求
// alitrip.flight.external.alipay.encrypt
//
// 支付宝小程序明文加密
type AlitripFlightExternalAlipayEncryptAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayEncryptReq *AlipayEncryptReq
}

// NewAlitripFlightExternalAlipayEncryptRequest 初始化AlitripFlightExternalAlipayEncryptAPIRequest对象
func NewAlitripFlightExternalAlipayEncryptRequest() *AlitripFlightExternalAlipayEncryptAPIRequest {
	return &AlitripFlightExternalAlipayEncryptAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightExternalAlipayEncryptAPIRequest) Reset() {
	r._alipayEncryptReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightExternalAlipayEncryptAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.encrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightExternalAlipayEncryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightExternalAlipayEncryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayEncryptReq is AlipayEncryptReq Setter
// 入参结构体
func (r *AlitripFlightExternalAlipayEncryptAPIRequest) SetAlipayEncryptReq(_alipayEncryptReq *AlipayEncryptReq) error {
	r._alipayEncryptReq = _alipayEncryptReq
	r.Set("alipay_encrypt_req", _alipayEncryptReq)
	return nil
}

// GetAlipayEncryptReq AlipayEncryptReq Getter
func (r AlitripFlightExternalAlipayEncryptAPIRequest) GetAlipayEncryptReq() *AlipayEncryptReq {
	return r._alipayEncryptReq
}

var poolAlitripFlightExternalAlipayEncryptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightExternalAlipayEncryptRequest()
	},
}

// GetAlitripFlightExternalAlipayEncryptRequest 从 sync.Pool 获取 AlitripFlightExternalAlipayEncryptAPIRequest
func GetAlitripFlightExternalAlipayEncryptAPIRequest() *AlitripFlightExternalAlipayEncryptAPIRequest {
	return poolAlitripFlightExternalAlipayEncryptAPIRequest.Get().(*AlitripFlightExternalAlipayEncryptAPIRequest)
}

// ReleaseAlitripFlightExternalAlipayEncryptAPIRequest 将 AlitripFlightExternalAlipayEncryptAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightExternalAlipayEncryptAPIRequest(v *AlitripFlightExternalAlipayEncryptAPIRequest) {
	v.Reset()
	poolAlitripFlightExternalAlipayEncryptAPIRequest.Put(v)
}
