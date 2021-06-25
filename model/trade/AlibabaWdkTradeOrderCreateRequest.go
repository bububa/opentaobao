package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部交易订单创单接口 APIRequest
alibaba.wdk.trade.order.create

通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderCreateRequest struct {
    model.Params

    // 待创建的订单
    trade   *TradeOrder 

}

func NewAlibabaWdkTradeOrderCreateRequest() *AlibabaWdkTradeOrderCreateRequest{
    return &AlibabaWdkTradeOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTradeOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.order.create"
}

func (r AlibabaWdkTradeOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTradeOrderCreateRequest) SetTrade(trade *TradeOrder) error {
    r.trade = trade
    r.Set("trade", trade)
    return nil
}

func (r AlibabaWdkTradeOrderCreateRequest) GetTrade() *TradeOrder {
    return r.trade
}

