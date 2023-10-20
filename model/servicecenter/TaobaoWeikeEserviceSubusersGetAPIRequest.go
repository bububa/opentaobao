package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweikeeservicesubusersgetAPIRequest 客服外包订单分配的商家子账号列表 API请求
// taobao.weike.eservice.subusers.get
//
// 获取客服外包订单分配的商家子账号列表，以及授权状态
type TaobaoweikeeservicesubusersgetAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// NewTaobaoweikeeservicesubusersgetRequest 初始化TaobaoweikeeservicesubusersgetAPIRequest对象
func NewTaobaoweikeeservicesubusersgetRequest() *TaobaoweikeeservicesubusersgetAPIRequest {
	return &TaobaoweikeeservicesubusersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweikeeservicesubusersgetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweikeeservicesubusersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweikeeservicesubusersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoweikeeservicesubusersgetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoweikeeservicesubusersgetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
