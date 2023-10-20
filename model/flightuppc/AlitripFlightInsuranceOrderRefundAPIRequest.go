package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceorderrefundAPIRequest 保险订单退保 API请求
// alitrip.flight.insurance.order.refund
//
// 保险订单退保
type AlitripflightinsuranceorderrefundAPIRequest struct {
	model.Params
	// 请求体
	_insRefundOrderReq *InsReverseOrderReq
}

// NewAlitripflightinsuranceorderrefundRequest 初始化AlitripflightinsuranceorderrefundAPIRequest对象
func NewAlitripflightinsuranceorderrefundRequest() *AlitripflightinsuranceorderrefundAPIRequest {
	return &AlitripflightinsuranceorderrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightinsuranceorderrefundAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightinsuranceorderrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightinsuranceorderrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsRefundOrderReq is InsRefundOrderReq Setter
// 请求体
func (r *AlitripflightinsuranceorderrefundAPIRequest) SetInsRefundOrderReq(_insRefundOrderReq *InsReverseOrderReq) error {
	r._insRefundOrderReq = _insRefundOrderReq
	r.Set("ins_refund_order_req", _insRefundOrderReq)
	return nil
}

// GetInsRefundOrderReq InsRefundOrderReq Getter
func (r AlitripflightinsuranceorderrefundAPIRequest) GetInsRefundOrderReq() *InsReverseOrderReq {
	return r._insRefundOrderReq
}
