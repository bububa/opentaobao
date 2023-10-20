package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceordermodifyAPIRequest 保险订单批改申请 API请求
// alitrip.flight.insurance.order.modify
//
// 保险订单批改申请
type AlitripflightinsuranceordermodifyAPIRequest struct {
	model.Params
	// 请求体
	_insReverseOrderReq *InsReverseOrderReq
}

// NewAlitripflightinsuranceordermodifyRequest 初始化AlitripflightinsuranceordermodifyAPIRequest对象
func NewAlitripflightinsuranceordermodifyRequest() *AlitripflightinsuranceordermodifyAPIRequest {
	return &AlitripflightinsuranceordermodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightinsuranceordermodifyAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightinsuranceordermodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightinsuranceordermodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInsReverseOrderReq is InsReverseOrderReq Setter
// 请求体
func (r *AlitripflightinsuranceordermodifyAPIRequest) SetInsReverseOrderReq(_insReverseOrderReq *InsReverseOrderReq) error {
	r._insReverseOrderReq = _insReverseOrderReq
	r.Set("ins_reverse_order_req", _insReverseOrderReq)
	return nil
}

// GetInsReverseOrderReq InsReverseOrderReq Getter
func (r AlitripflightinsuranceordermodifyAPIRequest) GetInsReverseOrderReq() *InsReverseOrderReq {
	return r._insReverseOrderReq
}
