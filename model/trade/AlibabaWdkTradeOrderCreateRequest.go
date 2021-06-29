package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部交易订单创单接口 API请求
alibaba.wdk.trade.order.create

通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderCreateRequest struct {
    model.Params
    // 待创建的订单
    _trade   *TradeOrder
}

// 初始化AlibabaWdkTradeOrderCreateRequest对象
func NewAlibabaWdkTradeOrderCreateRequest() *AlibabaWdkTradeOrderCreateRequest{
    return &AlibabaWdkTradeOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Trade Setter
// 待创建的订单
func (r *AlibabaWdkTradeOrderCreateRequest) SetTrade(_trade *TradeOrder) error {
    r._trade = _trade
    r.Set("trade", _trade)
    return nil
}

// Trade Getter
func (r AlibabaWdkTradeOrderCreateRequest) GetTrade() *TradeOrder {
    return r._trade
}
