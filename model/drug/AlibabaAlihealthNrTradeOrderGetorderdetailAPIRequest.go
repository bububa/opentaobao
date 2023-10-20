package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrtradeordergetorderdetailAPIRequest 根据订单id获取单条订单详情 API请求
// alibaba.alihealth.nr.trade.order.getorderdetail
//
// 阿里健康O2O，获取订单详情，修复组合商品价格精度问题
type AlibabaalihealthnrtradeordergetorderdetailAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// NewAlibabaalihealthnrtradeordergetorderdetailRequest 初始化AlibabaalihealthnrtradeordergetorderdetailAPIRequest对象
func NewAlibabaalihealthnrtradeordergetorderdetailRequest() *AlibabaalihealthnrtradeordergetorderdetailAPIRequest {
	return &AlibabaalihealthnrtradeordergetorderdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthnrtradeordergetorderdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.trade.order.getorderdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthnrtradeordergetorderdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthnrtradeordergetorderdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 淘宝订单ID
func (r *AlibabaalihealthnrtradeordergetorderdetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthnrtradeordergetorderdetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}
