package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部交易订单接口 API请求
alibaba.wdk.trade.order.query

通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderQueryAPIRequest struct {
    model.Params
    // 订单查询
    _query   *TradeOrderQuery
}

// 初始化AlibabaWdkTradeOrderQueryAPIRequest对象
func NewAlibabaWdkTradeOrderQueryRequest() *AlibabaWdkTradeOrderQueryAPIRequest{
    return &AlibabaWdkTradeOrderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 订单查询
func (r *AlibabaWdkTradeOrderQueryAPIRequest) SetQuery(_query *TradeOrderQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaWdkTradeOrderQueryAPIRequest) GetQuery() *TradeOrderQuery {
    return r._query
}
