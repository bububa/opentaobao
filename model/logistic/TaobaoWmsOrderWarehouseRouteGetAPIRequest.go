package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowmsorderwarehouseroutegetAPIRequest 获取订单仓库路由信息 API请求
// taobao.wms.order.warehouse.route.get
//
// 获取订单仓库路由信息
type TaobaowmsorderwarehouseroutegetAPIRequest struct {
	model.Params
	// 订单编号
	_orderCode string
}

// NewTaobaowmsorderwarehouseroutegetRequest 初始化TaobaowmsorderwarehouseroutegetAPIRequest对象
func NewTaobaowmsorderwarehouseroutegetRequest() *TaobaowmsorderwarehouseroutegetAPIRequest {
	return &TaobaowmsorderwarehouseroutegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowmsorderwarehouseroutegetAPIRequest) GetApiMethodName() string {
	return "taobao.wms.order.warehouse.route.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowmsorderwarehouseroutegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowmsorderwarehouseroutegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 订单编号
func (r *TaobaowmsorderwarehouseroutegetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowmsorderwarehouseroutegetAPIRequest) GetOrderCode() string {
	return r._orderCode
}
