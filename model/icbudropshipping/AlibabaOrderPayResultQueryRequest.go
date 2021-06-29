package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba查询订单支付结果 APIRequest
alibaba.order.pay.result.query

alibaba查询订单支付结果
*/
type AlibabaOrderPayResultQueryRequest struct {
    model.Params

    // order id
    tradeId   int64 

}

func NewAlibabaOrderPayResultQueryRequest() *AlibabaOrderPayResultQueryRequest{
    return &AlibabaOrderPayResultQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOrderPayResultQueryRequest) GetApiMethodName() string {
    return "alibaba.order.pay.result.query"
}

func (r AlibabaOrderPayResultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOrderPayResultQueryRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r AlibabaOrderPayResultQueryRequest) GetTradeId() int64 {
    return r.tradeId
}

