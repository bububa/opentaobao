package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWmsOrderWarehouseRouteGetAPIRequest 获取订单仓库路由信息 API请求
// taobao.wms.order.warehouse.route.get
//
// 获取订单仓库路由信息
type TaobaoWmsOrderWarehouseRouteGetAPIRequest struct {
	model.Params
	// 订单编号
	_orderCode string
}

// NewTaobaoWmsOrderWarehouseRouteGetRequest 初始化TaobaoWmsOrderWarehouseRouteGetAPIRequest对象
func NewTaobaoWmsOrderWarehouseRouteGetRequest() *TaobaoWmsOrderWarehouseRouteGetAPIRequest {
	return &TaobaoWmsOrderWarehouseRouteGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWmsOrderWarehouseRouteGetAPIRequest) Reset() {
	r._orderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWmsOrderWarehouseRouteGetAPIRequest) GetApiMethodName() string {
	return "taobao.wms.order.warehouse.route.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWmsOrderWarehouseRouteGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWmsOrderWarehouseRouteGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 订单编号
func (r *TaobaoWmsOrderWarehouseRouteGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWmsOrderWarehouseRouteGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

var poolTaobaoWmsOrderWarehouseRouteGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWmsOrderWarehouseRouteGetRequest()
	},
}

// GetTaobaoWmsOrderWarehouseRouteGetRequest 从 sync.Pool 获取 TaobaoWmsOrderWarehouseRouteGetAPIRequest
func GetTaobaoWmsOrderWarehouseRouteGetAPIRequest() *TaobaoWmsOrderWarehouseRouteGetAPIRequest {
	return poolTaobaoWmsOrderWarehouseRouteGetAPIRequest.Get().(*TaobaoWmsOrderWarehouseRouteGetAPIRequest)
}

// ReleaseTaobaoWmsOrderWarehouseRouteGetAPIRequest 将 TaobaoWmsOrderWarehouseRouteGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWmsOrderWarehouseRouteGetAPIRequest(v *TaobaoWmsOrderWarehouseRouteGetAPIRequest) {
	v.Reset()
	poolTaobaoWmsOrderWarehouseRouteGetAPIRequest.Put(v)
}
