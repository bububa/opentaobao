package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情 API请求
alibaba.happytrip.taxi.order.get

获取订单状态及详情
*/
type AlibabaHappytripTaxiOrderGetRequest struct {
    model.Params
    // 订单id
    orderId   string
}

// 初始化AlibabaHappytripTaxiOrderGetRequest对象
func NewAlibabaHappytripTaxiOrderGetRequest() *AlibabaHappytripTaxiOrderGetRequest{
    return &AlibabaHappytripTaxiOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderGetRequest) GetOrderId() string {
    return r.orderId
}
