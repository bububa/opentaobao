package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytripordergetAPIRequest 获取欢行统一订单模型 API请求
// alibaba.happytrip.order.get
//
// 通过订单id获取欢行统一订单模型数据
type AlibabahappytripordergetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabahappytripordergetRequest 初始化AlibabahappytripordergetAPIRequest对象
func NewAlibabahappytripordergetRequest() *AlibabahappytripordergetAPIRequest {
	return &AlibabahappytripordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytripordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytripordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytripordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahappytripordergetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytripordergetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
