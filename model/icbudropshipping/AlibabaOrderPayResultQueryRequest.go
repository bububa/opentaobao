package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba查询订单支付结果 API请求
alibaba.order.pay.result.query

alibaba查询订单支付结果
*/
type AlibabaOrderPayResultQueryRequest struct {
    model.Params
    // order id
    tradeId   int64
}

// 初始化AlibabaOrderPayResultQueryRequest对象
func NewAlibabaOrderPayResultQueryRequest() *AlibabaOrderPayResultQueryRequest{
    return &AlibabaOrderPayResultQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOrderPayResultQueryRequest) GetApiMethodName() string {
    return "alibaba.order.pay.result.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOrderPayResultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// order id
func (r *AlibabaOrderPayResultQueryRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r AlibabaOrderPayResultQueryRequest) GetTradeId() int64 {
    return r.tradeId
}
