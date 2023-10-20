package flightuppc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceOrderModifyAPIRequest 保险订单批改申请 API请求
// alitrip.flight.insurance.order.modify
//
// 保险订单批改申请
type AlitripFlightInsuranceOrderModifyAPIRequest struct {
	model.Params
	// 请求体
	_insReverseOrderReq *InsReverseOrderReq
}

// NewAlitripFlightInsuranceOrderModifyRequest 初始化AlitripFlightInsuranceOrderModifyAPIRequest对象
func NewAlitripFlightInsuranceOrderModifyRequest() *AlitripFlightInsuranceOrderModifyAPIRequest {
	return &AlitripFlightInsuranceOrderModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightInsuranceOrderModifyAPIRequest) Reset() {
	r._insReverseOrderReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightInsuranceOrderModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightInsuranceOrderModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightInsuranceOrderModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsReverseOrderReq is InsReverseOrderReq Setter
// 请求体
func (r *AlitripFlightInsuranceOrderModifyAPIRequest) SetInsReverseOrderReq(_insReverseOrderReq *InsReverseOrderReq) error {
	r._insReverseOrderReq = _insReverseOrderReq
	r.Set("ins_reverse_order_req", _insReverseOrderReq)
	return nil
}

// GetInsReverseOrderReq InsReverseOrderReq Getter
func (r AlitripFlightInsuranceOrderModifyAPIRequest) GetInsReverseOrderReq() *InsReverseOrderReq {
	return r._insReverseOrderReq
}

var poolAlitripFlightInsuranceOrderModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightInsuranceOrderModifyRequest()
	},
}

// GetAlitripFlightInsuranceOrderModifyRequest 从 sync.Pool 获取 AlitripFlightInsuranceOrderModifyAPIRequest
func GetAlitripFlightInsuranceOrderModifyAPIRequest() *AlitripFlightInsuranceOrderModifyAPIRequest {
	return poolAlitripFlightInsuranceOrderModifyAPIRequest.Get().(*AlitripFlightInsuranceOrderModifyAPIRequest)
}

// ReleaseAlitripFlightInsuranceOrderModifyAPIRequest 将 AlitripFlightInsuranceOrderModifyAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightInsuranceOrderModifyAPIRequest(v *AlitripFlightInsuranceOrderModifyAPIRequest) {
	v.Reset()
	poolAlitripFlightInsuranceOrderModifyAPIRequest.Put(v)
}
