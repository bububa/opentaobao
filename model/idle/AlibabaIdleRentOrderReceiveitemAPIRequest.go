package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentorderreceiveitemAPIRequest 确认签收 API请求
// alibaba.idle.rent.order.receiveitem
//
// 确认揽收/签收
type AlibabaidlerentorderreceiveitemAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabaidlerentorderreceiveitemRequest 初始化AlibabaidlerentorderreceiveitemAPIRequest对象
func NewAlibabaidlerentorderreceiveitemRequest() *AlibabaidlerentorderreceiveitemAPIRequest {
	return &AlibabaidlerentorderreceiveitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentorderreceiveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.receiveitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentorderreceiveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentorderreceiveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaidlerentorderreceiveitemAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaidlerentorderreceiveitemAPIRequest) GetOrderId() int64 {
	return r._orderId
}
