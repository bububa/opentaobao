package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriprentcarorderdetailqueryAPIRequest 租车订单详情查询 API请求
// alitrip.rentcar.order.detail.query
//
// 租车订单详情查询
type AlitriprentcarorderdetailqueryAPIRequest struct {
	model.Params
	// RQ
	_paramRentCarOrderDetailCallbackReq *RentCarOrderDetailCallbackReq
}

// NewAlitriprentcarorderdetailqueryRequest 初始化AlitriprentcarorderdetailqueryAPIRequest对象
func NewAlitriprentcarorderdetailqueryRequest() *AlitriprentcarorderdetailqueryAPIRequest {
	return &AlitriprentcarorderdetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriprentcarorderdetailqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.rentcar.order.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriprentcarorderdetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriprentcarorderdetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRentCarOrderDetailCallbackReq is ParamRentCarOrderDetailCallbackReq Setter
// RQ
func (r *AlitriprentcarorderdetailqueryAPIRequest) SetParamRentCarOrderDetailCallbackReq(_paramRentCarOrderDetailCallbackReq *RentCarOrderDetailCallbackReq) error {
	r._paramRentCarOrderDetailCallbackReq = _paramRentCarOrderDetailCallbackReq
	r.Set("param_rent_car_order_detail_callback_req", _paramRentCarOrderDetailCallbackReq)
	return nil
}

// GetParamRentCarOrderDetailCallbackReq ParamRentCarOrderDetailCallbackReq Getter
func (r AlitriprentcarorderdetailqueryAPIRequest) GetParamRentCarOrderDetailCallbackReq() *RentCarOrderDetailCallbackReq {
	return r._paramRentCarOrderDetailCallbackReq
}
