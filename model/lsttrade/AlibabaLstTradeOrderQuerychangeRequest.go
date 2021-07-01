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
type AlibabaLstTradeOrderQuerychangeAPIRequest struct {
    model.Params
    // 入参包装类
    _query   *LstOrderQuery
}

// 初始化AlibabaLstTradeOrderQuerychangeAPIRequest对象
func NewAlibabaLstTradeOrderQuerychangeRequest() *AlibabaLstTradeOrderQuerychangeAPIRequest{
    return &AlibabaLstTradeOrderQuerychangeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderQuerychangeAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.querychange"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderQuerychangeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参包装类
func (r *AlibabaLstTradeOrderQuerychangeAPIRequest) SetQuery(_query *LstOrderQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaLstTradeOrderQuerychangeAPIRequest) GetQuery() *LstOrderQuery {
    return r._query
}
