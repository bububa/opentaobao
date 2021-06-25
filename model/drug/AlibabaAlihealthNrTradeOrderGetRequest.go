package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单详情 APIRequest
alibaba.alihealth.nr.trade.order.get

阿里健康O2O，获取订单详情
*/
type AlibabaAlihealthNrTradeOrderGetRequest struct {
    model.Params

    // 淘宝订单ID
    orderId   int64 

}

func NewAlibabaAlihealthNrTradeOrderGetRequest() *AlibabaAlihealthNrTradeOrderGetRequest{
    return &AlibabaAlihealthNrTradeOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthNrTradeOrderGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.trade.order.get"
}

func (r AlibabaAlihealthNrTradeOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthNrTradeOrderGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthNrTradeOrderGetRequest) GetOrderId() int64 {
    return r.orderId
}

