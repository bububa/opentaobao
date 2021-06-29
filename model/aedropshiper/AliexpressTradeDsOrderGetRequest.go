package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
买家查询订单详情 API请求
aliexpress.trade.ds.order.get

买家查询订单详情，用于dropshipper
*/
type AliexpressTradeDsOrderGetRequest struct {
    model.Params
    // 订单查询条件
    _singleOrderQuery   *AeopSingleOrderQuery
}

// 初始化AliexpressTradeDsOrderGetRequest对象
func NewAliexpressTradeDsOrderGetRequest() *AliexpressTradeDsOrderGetRequest{
    return &AliexpressTradeDsOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTradeDsOrderGetRequest) GetApiMethodName() string {
    return "aliexpress.trade.ds.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTradeDsOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SingleOrderQuery Setter
// 订单查询条件
func (r *AliexpressTradeDsOrderGetRequest) SetSingleOrderQuery(_singleOrderQuery *AeopSingleOrderQuery) error {
    r._singleOrderQuery = _singleOrderQuery
    r.Set("single_order_query", _singleOrderQuery)
    return nil
}

// SingleOrderQuery Getter
func (r AliexpressTradeDsOrderGetRequest) GetSingleOrderQuery() *AeopSingleOrderQuery {
    return r._singleOrderQuery
}
