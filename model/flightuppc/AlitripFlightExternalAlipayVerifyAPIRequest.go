package flightuppc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightExternalAlipayVerifyAPIRequest 支付宝小程序验签 API请求
// alitrip.flight.external.alipay.verify
//
// 支付宝小程序验签
type AlitripFlightExternalAlipayVerifyAPIRequest struct {
	model.Params
	// 入参结构体
	_alipayCheckSignReq *AlipayCheckSignReq
}

// NewAlitripFlightExternalAlipayVerifyRequest 初始化AlitripFlightExternalAlipayVerifyAPIRequest对象
func NewAlitripFlightExternalAlipayVerifyRequest() *AlitripFlightExternalAlipayVerifyAPIRequest {
	return &AlitripFlightExternalAlipayVerifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightExternalAlipayVerifyAPIRequest) Reset() {
	r._alipayCheckSignReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightExternalAlipayVerifyAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.external.alipay.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightExternalAlipayVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightExternalAlipayVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayCheckSignReq is AlipayCheckSignReq Setter
// 入参结构体
func (r *AlitripFlightExternalAlipayVerifyAPIRequest) SetAlipayCheckSignReq(_alipayCheckSignReq *AlipayCheckSignReq) error {
	r._alipayCheckSignReq = _alipayCheckSignReq
	r.Set("alipay_check_sign_req", _alipayCheckSignReq)
	return nil
}

// GetAlipayCheckSignReq AlipayCheckSignReq Getter
func (r AlitripFlightExternalAlipayVerifyAPIRequest) GetAlipayCheckSignReq() *AlipayCheckSignReq {
	return r._alipayCheckSignReq
}

var poolAlitripFlightExternalAlipayVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightExternalAlipayVerifyRequest()
	},
}

// GetAlitripFlightExternalAlipayVerifyRequest 从 sync.Pool 获取 AlitripFlightExternalAlipayVerifyAPIRequest
func GetAlitripFlightExternalAlipayVerifyAPIRequest() *AlitripFlightExternalAlipayVerifyAPIRequest {
	return poolAlitripFlightExternalAlipayVerifyAPIRequest.Get().(*AlitripFlightExternalAlipayVerifyAPIRequest)
}

// ReleaseAlitripFlightExternalAlipayVerifyAPIRequest 将 AlitripFlightExternalAlipayVerifyAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightExternalAlipayVerifyAPIRequest(v *AlitripFlightExternalAlipayVerifyAPIRequest) {
	v.Reset()
	poolAlitripFlightExternalAlipayVerifyAPIRequest.Put(v)
}
