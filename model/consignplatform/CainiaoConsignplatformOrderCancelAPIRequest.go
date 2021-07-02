package consignplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoConsignplatformOrderCancelAPIRequest 菜鸟发货工作台取消包裹以及订单 API请求
// cainiao.consignplatform.order.cancel
//
// 菜鸟发货工作台，商家或者isv通过api取消包裹、回收单号，如果是裹裹运力会取消小件员上门。最后删除订单信息。
type CainiaoConsignplatformOrderCancelAPIRequest struct {
	model.Params
	// 取消参数
	_cancelRequest *OrderCancelRequest
}

// NewCainiaoConsignplatformOrderCancelRequest 初始化CainiaoConsignplatformOrderCancelAPIRequest对象
func NewCainiaoConsignplatformOrderCancelRequest() *CainiaoConsignplatformOrderCancelAPIRequest {
	return &CainiaoConsignplatformOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoConsignplatformOrderCancelAPIRequest) GetApiMethodName() string {
	return "cainiao.consignplatform.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoConsignplatformOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CancelRequest Setter
// 取消参数
func (r *CainiaoConsignplatformOrderCancelAPIRequest) SetCancelRequest(_cancelRequest *OrderCancelRequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// Get CancelRequest Getter
func (r CainiaoConsignplatformOrderCancelAPIRequest) GetCancelRequest() *OrderCancelRequest {
	return r._cancelRequest
}
