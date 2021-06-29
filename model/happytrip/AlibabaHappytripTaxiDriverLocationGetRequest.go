package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
司机位置 APIRequest
alibaba.happytrip.taxi.driver.location.get

获取司机实时位置
*/
type AlibabaHappytripTaxiDriverLocationGetRequest struct {
    model.Params

    // 供应商订单号
    orderId   string 

}

func NewAlibabaHappytripTaxiDriverLocationGetRequest() *AlibabaHappytripTaxiDriverLocationGetRequest{
    return &AlibabaHappytripTaxiDriverLocationGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiDriverLocationGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.driver.location.get"
}

func (r AlibabaHappytripTaxiDriverLocationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiDriverLocationGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiDriverLocationGetRequest) GetOrderId() string {
    return r.orderId
}

