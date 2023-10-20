package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrtradeordergetAPIRequest 获取订单详情 API请求
// alibaba.alihealth.nr.trade.order.get
//
// 阿里健康O2O，获取订单详情
type AlibabaalihealthnrtradeordergetAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// NewAlibabaalihealthnrtradeordergetRequest 初始化AlibabaalihealthnrtradeordergetAPIRequest对象
func NewAlibabaalihealthnrtradeordergetRequest() *AlibabaalihealthnrtradeordergetAPIRequest {
	return &AlibabaalihealthnrtradeordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthnrtradeordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.trade.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthnrtradeordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthnrtradeordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 淘宝订单ID
func (r *AlibabaalihealthnrtradeordergetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthnrtradeordergetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
