package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情查询接口 APIRequest
alibaba.health.nr.cep.order.query

订单详情查询接口
*/
type AlibabaHealthNrCepOrderQueryRequest struct {
    model.Params

    // 订单号
    orderId   int64 

}

func NewAlibabaHealthNrCepOrderQueryRequest() *AlibabaHealthNrCepOrderQueryRequest{
    return &AlibabaHealthNrCepOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthNrCepOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.order.query"
}

func (r AlibabaHealthNrCepOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthNrCepOrderQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthNrCepOrderQueryRequest) GetOrderId() int64 {
    return r.orderId
}

