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
type AlibabaOrderPayResultQueryAPIRequest struct {
    model.Params
    // order id
    _tradeId   int64
}

// 初始化AlibabaOrderPayResultQueryAPIRequest对象
func NewAlibabaOrderPayResultQueryRequest() *AlibabaOrderPayResultQueryAPIRequest{
    return &AlibabaOrderPayResultQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOrderPayResultQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.order.pay.result.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOrderPayResultQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// order id
func (r *AlibabaOrderPayResultQueryAPIRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r AlibabaOrderPayResultQueryAPIRequest) GetTradeId() int64 {
    return r._tradeId
}
