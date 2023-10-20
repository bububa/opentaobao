package flightuppc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceOrderRefundAPIRequest 保险订单退保 API请求
// alitrip.flight.insurance.order.refund
//
// 保险订单退保
type AlitripFlightInsuranceOrderRefundAPIRequest struct {
	model.Params
	// 请求体
	_insRefundOrderReq *InsReverseOrderReq
}

// NewAlitripFlightInsuranceOrderRefundRequest 初始化AlitripFlightInsuranceOrderRefundAPIRequest对象
func NewAlitripFlightInsuranceOrderRefundRequest() *AlitripFlightInsuranceOrderRefundAPIRequest {
	return &AlitripFlightInsuranceOrderRefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightInsuranceOrderRefundAPIRequest) Reset() {
	r._insRefundOrderReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightInsuranceOrderRefundAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightInsuranceOrderRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightInsuranceOrderRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsRefundOrderReq is InsRefundOrderReq Setter
// 请求体
func (r *AlitripFlightInsuranceOrderRefundAPIRequest) SetInsRefundOrderReq(_insRefundOrderReq *InsReverseOrderReq) error {
	r._insRefundOrderReq = _insRefundOrderReq
	r.Set("ins_refund_order_req", _insRefundOrderReq)
	return nil
}

// GetInsRefundOrderReq InsRefundOrderReq Getter
func (r AlitripFlightInsuranceOrderRefundAPIRequest) GetInsRefundOrderReq() *InsReverseOrderReq {
	return r._insRefundOrderReq
}

var poolAlitripFlightInsuranceOrderRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightInsuranceOrderRefundRequest()
	},
}

// GetAlitripFlightInsuranceOrderRefundRequest 从 sync.Pool 获取 AlitripFlightInsuranceOrderRefundAPIRequest
func GetAlitripFlightInsuranceOrderRefundAPIRequest() *AlitripFlightInsuranceOrderRefundAPIRequest {
	return poolAlitripFlightInsuranceOrderRefundAPIRequest.Get().(*AlitripFlightInsuranceOrderRefundAPIRequest)
}

// ReleaseAlitripFlightInsuranceOrderRefundAPIRequest 将 AlitripFlightInsuranceOrderRefundAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightInsuranceOrderRefundAPIRequest(v *AlitripFlightInsuranceOrderRefundAPIRequest) {
	v.Reset()
	poolAlitripFlightInsuranceOrderRefundAPIRequest.Put(v)
}
