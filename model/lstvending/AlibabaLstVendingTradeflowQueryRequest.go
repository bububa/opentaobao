package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机交易流水查询 API请求
alibaba.lst.vending.tradeflow.query

零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。
*/
type AlibabaLstVendingTradeflowQueryAPIRequest struct {
    model.Params
    // 交易流水查询条件
    _openTradeFlowQuery   *OpenTradeFlowQuery
}

// 初始化AlibabaLstVendingTradeflowQueryAPIRequest对象
func NewAlibabaLstVendingTradeflowQueryRequest() *AlibabaLstVendingTradeflowQueryAPIRequest{
    return &AlibabaLstVendingTradeflowQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingTradeflowQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.tradeflow.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingTradeflowQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenTradeFlowQuery Setter
// 交易流水查询条件
func (r *AlibabaLstVendingTradeflowQueryAPIRequest) SetOpenTradeFlowQuery(_openTradeFlowQuery *OpenTradeFlowQuery) error {
    r._openTradeFlowQuery = _openTradeFlowQuery
    r.Set("open_trade_flow_query", _openTradeFlowQuery)
    return nil
}

// OpenTradeFlowQuery Getter
func (r AlibabaLstVendingTradeflowQueryAPIRequest) GetOpenTradeFlowQuery() *OpenTradeFlowQuery {
    return r._openTradeFlowQuery
}
