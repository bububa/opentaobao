package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderConfirmAPIRequest 费用确认 API请求
// alibaba.happytrip.taxi.order.confirm
//
// 1.司机点结束计费,欢行会收到正常支付待评论 回调,确认费用无误欢行可以通过此接口确认并支付。
// 2.如果欢行一直不调用此接口,订单会在48小时后自动支付。
type AlibabaHappytripTaxiOrderConfirmAPIRequest struct {
	model.Params
	// 要确认支付的订单号
	_orderId string
}

// NewAlibabaHappytripTaxiOrderConfirmRequest 初始化AlibabaHappytripTaxiOrderConfirmAPIRequest对象
func NewAlibabaHappytripTaxiOrderConfirmRequest() *AlibabaHappytripTaxiOrderConfirmAPIRequest {
	return &AlibabaHappytripTaxiOrderConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 要确认支付的订单号
func (r *AlibabaHappytripTaxiOrderConfirmAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiOrderConfirmAPIRequest) GetOrderId() string {
	return r._orderId
}
