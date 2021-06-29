package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单 APIRequest
alibaba.idle.rent.order.query

查询订单信息
*/
type AlibabaIdleRentOrderQueryRequest struct {
    model.Params

    // 订单id
    orderId   int64 

}

func NewAlibabaIdleRentOrderQueryRequest() *AlibabaIdleRentOrderQueryRequest{
    return &AlibabaIdleRentOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.query"
}

func (r AlibabaIdleRentOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentOrderQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaIdleRentOrderQueryRequest) GetOrderId() int64 {
    return r.orderId
}

