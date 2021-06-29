package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单指派 APIRequest
alibaba.happytrip.taxi.order.assign

通知供应商订单指派成功
*/
type AlibabaHappytripTaxiOrderAssignRequest struct {
    model.Params

    // 供应商订单号
    orderId   string 

}

func NewAlibabaHappytripTaxiOrderAssignRequest() *AlibabaHappytripTaxiOrderAssignRequest{
    return &AlibabaHappytripTaxiOrderAssignRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderAssignRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.assign"
}

func (r AlibabaHappytripTaxiOrderAssignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderAssignRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderAssignRequest) GetOrderId() string {
    return r.orderId
}

