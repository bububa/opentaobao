package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单 APIRequest
youku.ott.pay.order.queryorder

通过订单号查询订单信息
*/
type YoukuOttPayOrderQueryorderRequest struct {
    model.Params

    // 订单号
    orderNo   string 

}

func NewYoukuOttPayOrderQueryorderRequest() *YoukuOttPayOrderQueryorderRequest{
    return &YoukuOttPayOrderQueryorderRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPayOrderQueryorderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.queryorder"
}

func (r YoukuOttPayOrderQueryorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPayOrderQueryorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r YoukuOttPayOrderQueryorderRequest) GetOrderNo() string {
    return r.orderNo
}

