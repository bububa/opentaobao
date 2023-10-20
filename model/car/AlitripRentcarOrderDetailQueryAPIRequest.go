package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripRentcarOrderDetailQueryAPIRequest 租车订单详情查询 API请求
// alitrip.rentcar.order.detail.query
//
// 租车订单详情查询
type AlitripRentcarOrderDetailQueryAPIRequest struct {
	model.Params
	// RQ
	_paramRentCarOrderDetailCallbackReq *RentCarOrderDetailCallbackReq
}

// NewAlitripRentcarOrderDetailQueryRequest 初始化AlitripRentcarOrderDetailQueryAPIRequest对象
func NewAlitripRentcarOrderDetailQueryRequest() *AlitripRentcarOrderDetailQueryAPIRequest {
	return &AlitripRentcarOrderDetailQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRentcarOrderDetailQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.rentcar.order.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRentcarOrderDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripRentcarOrderDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRentCarOrderDetailCallbackReq is ParamRentCarOrderDetailCallbackReq Setter
// RQ
func (r *AlitripRentcarOrderDetailQueryAPIRequest) SetParamRentCarOrderDetailCallbackReq(_paramRentCarOrderDetailCallbackReq *RentCarOrderDetailCallbackReq) error {
	r._paramRentCarOrderDetailCallbackReq = _paramRentCarOrderDetailCallbackReq
	r.Set("param_rent_car_order_detail_callback_req", _paramRentCarOrderDetailCallbackReq)
	return nil
}

// GetParamRentCarOrderDetailCallbackReq ParamRentCarOrderDetailCallbackReq Getter
func (r AlitripRentcarOrderDetailQueryAPIRequest) GetParamRentCarOrderDetailCallbackReq() *RentCarOrderDetailCallbackReq {
	return r._paramRentCarOrderDetailCallbackReq
}
