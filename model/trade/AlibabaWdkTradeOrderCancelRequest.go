package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部交易订单取消接口 APIRequest
alibaba.wdk.trade.order.cancel

通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderCancelRequest struct {
    model.Params

    // 待取消的订单
    trade   *TradeOrder 

}

func NewAlibabaWdkTradeOrderCancelRequest() *AlibabaWdkTradeOrderCancelRequest{
    return &AlibabaWdkTradeOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTradeOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.order.cancel"
}

func (r AlibabaWdkTradeOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTradeOrderCancelRequest) SetTrade(trade *TradeOrder) error {
    r.trade = trade
    r.Set("trade", trade)
    return nil
}

func (r AlibabaWdkTradeOrderCancelRequest) GetTrade() *TradeOrder {
    return r.trade
}

