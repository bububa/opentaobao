package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeDrugOrderGetAPIRequest
查看订单详情 API请求
taobao.trade.drug.order.get

商家查看订单详情 */
type TaobaoTradeDrugOrderGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewTaobaoTradeDrugOrderGetRequest 初始化TaobaoTradeDrugOrderGetAPIRequest对象
func NewTaobaoTradeDrugOrderGetRequest() *TaobaoTradeDrugOrderGetAPIRequest {
	return &TaobaoTradeDrugOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *TaobaoTradeDrugOrderGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoTradeDrugOrderGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
