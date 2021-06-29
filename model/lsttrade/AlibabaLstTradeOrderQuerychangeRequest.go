package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单id批量查询（品牌商视角） APIRequest
alibaba.lst.trade.order.querychange

根据品牌和时间段查询有变更记录的订单id
*/
type AlibabaLstTradeOrderQuerychangeRequest struct {
    model.Params

    // 入参包装类
    query   *LstOrderQuery 

}

func NewAlibabaLstTradeOrderQuerychangeRequest() *AlibabaLstTradeOrderQuerychangeRequest{
    return &AlibabaLstTradeOrderQuerychangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeOrderQuerychangeRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.querychange"
}

func (r AlibabaLstTradeOrderQuerychangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeOrderQuerychangeRequest) SetQuery(query *LstOrderQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaLstTradeOrderQuerychangeRequest) GetQuery() *LstOrderQuery {
    return r.query
}

