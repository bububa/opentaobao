package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单id批量查询（品牌商视角） API请求
alibaba.lst.trade.order.querychange

根据品牌和时间段查询有变更记录的订单id
*/
type AlibabaLstTradeOrderQuerychangeRequest struct {
    model.Params
    // 入参包装类
    query   *LstOrderQuery
}

// 初始化AlibabaLstTradeOrderQuerychangeRequest对象
func NewAlibabaLstTradeOrderQuerychangeRequest() *AlibabaLstTradeOrderQuerychangeRequest{
    return &AlibabaLstTradeOrderQuerychangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderQuerychangeRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.querychange"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderQuerychangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参包装类
func (r *AlibabaLstTradeOrderQuerychangeRequest) SetQuery(query *LstOrderQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaLstTradeOrderQuerychangeRequest) GetQuery() *LstOrderQuery {
    return r.query
}
