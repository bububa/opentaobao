package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
费用确认 API请求
alibaba.happytrip.taxi.order.confirm

1.司机点结束计费,欢行会收到正常支付待评论 回调,确认费用无误欢行可以通过此接口确认并支付。
2.如果欢行一直不调用此接口,订单会在48小时后自动支付。
*/
type AlibabaHappytripTaxiOrderConfirmRequest struct {
    model.Params
    // 要确认支付的订单号
    orderId   string
}

// 初始化AlibabaHappytripTaxiOrderConfirmRequest对象
func NewAlibabaHappytripTaxiOrderConfirmRequest() *AlibabaHappytripTaxiOrderConfirmRequest{
    return &AlibabaHappytripTaxiOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderConfirmRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 要确认支付的订单号
func (r *AlibabaHappytripTaxiOrderConfirmRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderConfirmRequest) GetOrderId() string {
    return r.orderId
}
