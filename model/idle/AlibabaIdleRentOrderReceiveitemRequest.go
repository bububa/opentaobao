package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认签收 APIRequest
alibaba.idle.rent.order.receiveitem

确认揽收/签收
*/
type AlibabaIdleRentOrderReceiveitemRequest struct {
    model.Params

    // 订单id
    orderId   int64 

}

func NewAlibabaIdleRentOrderReceiveitemRequest() *AlibabaIdleRentOrderReceiveitemRequest{
    return &AlibabaIdleRentOrderReceiveitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentOrderReceiveitemRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.receiveitem"
}

func (r AlibabaIdleRentOrderReceiveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentOrderReceiveitemRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaIdleRentOrderReceiveitemRequest) GetOrderId() int64 {
    return r.orderId
}

