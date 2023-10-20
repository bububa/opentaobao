package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUopTobOrderCreateAPIRequest ToB仓储发货 API请求
// taobao.uop.tob.order.create
//
// ToB仓储发货
type TaobaoUopTobOrderCreateAPIRequest struct {
	model.Params
	// ERP出库对象
	_deliveryOrder *DeliveryOrder
}

// NewTaobaoUopTobOrderCreateRequest 初始化TaobaoUopTobOrderCreateAPIRequest对象
func NewTaobaoUopTobOrderCreateRequest() *TaobaoUopTobOrderCreateAPIRequest {
	return &TaobaoUopTobOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUopTobOrderCreateAPIRequest) Reset() {
	r._deliveryOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUopTobOrderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.uop.tob.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUopTobOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUopTobOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryOrder is DeliveryOrder Setter
// ERP出库对象
func (r *TaobaoUopTobOrderCreateAPIRequest) SetDeliveryOrder(_deliveryOrder *DeliveryOrder) error {
	r._deliveryOrder = _deliveryOrder
	r.Set("delivery_order", _deliveryOrder)
	return nil
}

// GetDeliveryOrder DeliveryOrder Getter
func (r TaobaoUopTobOrderCreateAPIRequest) GetDeliveryOrder() *DeliveryOrder {
	return r._deliveryOrder
}

var poolTaobaoUopTobOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUopTobOrderCreateRequest()
	},
}

// GetTaobaoUopTobOrderCreateRequest 从 sync.Pool 获取 TaobaoUopTobOrderCreateAPIRequest
func GetTaobaoUopTobOrderCreateAPIRequest() *TaobaoUopTobOrderCreateAPIRequest {
	return poolTaobaoUopTobOrderCreateAPIRequest.Get().(*TaobaoUopTobOrderCreateAPIRequest)
}

// ReleaseTaobaoUopTobOrderCreateAPIRequest 将 TaobaoUopTobOrderCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoUopTobOrderCreateAPIRequest(v *TaobaoUopTobOrderCreateAPIRequest) {
	v.Reset()
	poolTaobaoUopTobOrderCreateAPIRequest.Put(v)
}
