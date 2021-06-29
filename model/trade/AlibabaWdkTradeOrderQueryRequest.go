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
type AlibabaWdkTradeOrderQueryRequest struct {
    model.Params
    // 订单查询
    _query   *TradeOrderQuery
}

// 初始化AlibabaWdkTradeOrderQueryRequest对象
func NewAlibabaWdkTradeOrderQueryRequest() *AlibabaWdkTradeOrderQueryRequest{
    return &AlibabaWdkTradeOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 订单查询
func (r *AlibabaWdkTradeOrderQueryRequest) SetQuery(_query *TradeOrderQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaWdkTradeOrderQueryRequest) GetQuery() *TradeOrderQuery {
    return r._query
}
