package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消叫车 APIRequest
alibaba.happytrip.taxi.order.cancel

取消叫车订单,行程中的订单不能取消
*/
type AlibabaHappytripTaxiOrderCancelRequest struct {
    model.Params

    // 订单id
    orderId   string 

    // 是否强制取消(true 或 false)默认false
    force   string 

    // 取消类型，0：系统取消，非0：用户取消
    type   int64 

}

func NewAlibabaHappytripTaxiOrderCancelRequest() *AlibabaHappytripTaxiOrderCancelRequest{
    return &AlibabaHappytripTaxiOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.cancel"
}

func (r AlibabaHappytripTaxiOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderCancelRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderCancelRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHappytripTaxiOrderCancelRequest) SetForce(force string) error {
    r.force = force
    r.Set("force", force)
    return nil
}

func (r AlibabaHappytripTaxiOrderCancelRequest) GetForce() string {
    return r.force
}

func (r *AlibabaHappytripTaxiOrderCancelRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaHappytripTaxiOrderCancelRequest) GetType() int64 {
    return r.type
}

