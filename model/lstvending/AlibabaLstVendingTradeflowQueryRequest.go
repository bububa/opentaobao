package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机交易流水查询 APIRequest
alibaba.lst.vending.tradeflow.query

零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。
*/
type AlibabaLstVendingTradeflowQueryRequest struct {
    model.Params

    // 交易流水查询条件
    openTradeFlowQuery   *OpenTradeFlowQuery 

}

func NewAlibabaLstVendingTradeflowQueryRequest() *AlibabaLstVendingTradeflowQueryRequest{
    return &AlibabaLstVendingTradeflowQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstVendingTradeflowQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.tradeflow.query"
}

func (r AlibabaLstVendingTradeflowQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstVendingTradeflowQueryRequest) SetOpenTradeFlowQuery(openTradeFlowQuery *OpenTradeFlowQuery) error {
    r.openTradeFlowQuery = openTradeFlowQuery
    r.Set("open_trade_flow_query", openTradeFlowQuery)
    return nil
}

func (r AlibabaLstVendingTradeflowQueryRequest) GetOpenTradeFlowQuery() *OpenTradeFlowQuery {
    return r.openTradeFlowQuery
}

