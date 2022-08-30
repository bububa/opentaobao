package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusMerchantOrderGetAPIRequest 商家侧查询订单详情 API请求
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
type TaobaoBusMerchantOrderGetAPIRequest struct {
	model.Params
	// 飞猪订单号或商家订单号，多个以“,”分割
	_orderId string
}

// NewTaobaoBusMerchantOrderGetRequest 初始化TaobaoBusMerchantOrderGetAPIRequest对象
func NewTaobaoBusMerchantOrderGetRequest() *TaobaoBusMerchantOrderGetAPIRequest {
	return &TaobaoBusMerchantOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusMerchantOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.merchant.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusMerchantOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 飞猪订单号或商家订单号，多个以“,”分割
func (r *TaobaoBusMerchantOrderGetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoBusMerchantOrderGetAPIRequest) GetOrderId() string {
	return r._orderId
}
