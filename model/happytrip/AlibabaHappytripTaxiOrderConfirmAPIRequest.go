package happytrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiOrderConfirmAPIRequest) Reset() {
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiOrderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaHappytripTaxiOrderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiOrderConfirmRequest()
	},
}

// GetAlibabaHappytripTaxiOrderConfirmRequest 从 sync.Pool 获取 AlibabaHappytripTaxiOrderConfirmAPIRequest
func GetAlibabaHappytripTaxiOrderConfirmAPIRequest() *AlibabaHappytripTaxiOrderConfirmAPIRequest {
	return poolAlibabaHappytripTaxiOrderConfirmAPIRequest.Get().(*AlibabaHappytripTaxiOrderConfirmAPIRequest)
}

// ReleaseAlibabaHappytripTaxiOrderConfirmAPIRequest 将 AlibabaHappytripTaxiOrderConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderConfirmAPIRequest(v *AlibabaHappytripTaxiOrderConfirmAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderConfirmAPIRequest.Put(v)
}
