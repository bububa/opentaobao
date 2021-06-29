package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情 APIRequest
alibaba.happytrip.taxi.order.get

获取订单状态及详情
*/
type AlibabaHappytripTaxiOrderGetRequest struct {
    model.Params

    // 订单id
    orderId   string 

}

func NewAlibabaHappytripTaxiOrderGetRequest() *AlibabaHappytripTaxiOrderGetRequest{
    return &AlibabaHappytripTaxiOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.get"
}

func (r AlibabaHappytripTaxiOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderGetRequest) GetOrderId() string {
    return r.orderId
}

