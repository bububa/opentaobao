package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouoptobordercreateAPIRequest ToB仓储发货 API请求
// taobao.uop.tob.order.create
//
// ToB仓储发货
type TaobaouoptobordercreateAPIRequest struct {
	model.Params
	// ERP出库对象
	_deliveryOrder *DeliveryOrder
}

// NewTaobaouoptobordercreateRequest 初始化TaobaouoptobordercreateAPIRequest对象
func NewTaobaouoptobordercreateRequest() *TaobaouoptobordercreateAPIRequest {
	return &TaobaouoptobordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouoptobordercreateAPIRequest) GetApiMethodName() string {
	return "taobao.uop.tob.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouoptobordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouoptobordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryOrder is DeliveryOrder Setter
// ERP出库对象
func (r *TaobaouoptobordercreateAPIRequest) SetDeliveryOrder(_deliveryOrder *DeliveryOrder) error {
	r._deliveryOrder = _deliveryOrder
	r.Set("delivery_order", _deliveryOrder)
	return nil
}

// GetDeliveryOrder DeliveryOrder Getter
func (r TaobaouoptobordercreateAPIRequest) GetDeliveryOrder() *DeliveryOrder {
	return r._deliveryOrder
}
