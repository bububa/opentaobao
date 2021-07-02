package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest 根据订单id获取单条订单详情 API请求
// alibaba.alihealth.nr.trade.order.getorderdetail
//
// 阿里健康O2O，获取订单详情，修复组合商品价格精度问题
type AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// NewAlibabaAlihealthNrTradeOrderGetorderdetailRequest 初始化AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest对象
func NewAlibabaAlihealthNrTradeOrderGetorderdetailRequest() *AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest {
	return &AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.trade.order.getorderdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 淘宝订单ID
func (r *AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}
