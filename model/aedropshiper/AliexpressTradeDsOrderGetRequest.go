package aedropshiper

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
买家查询订单详情 APIRequest
aliexpress.trade.ds.order.get

买家查询订单详情，用于dropshipper
*/
type AliexpressTradeDsOrderGetRequest struct {
    model.Params

    // 订单查询条件
    singleOrderQuery   *AeopSingleOrderQuery 

}

func NewAliexpressTradeDsOrderGetRequest() *AliexpressTradeDsOrderGetRequest{
    return &AliexpressTradeDsOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressTradeDsOrderGetRequest) GetApiMethodName() string {
    return "aliexpress.trade.ds.order.get"
}

func (r AliexpressTradeDsOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressTradeDsOrderGetRequest) SetSingleOrderQuery(singleOrderQuery *AeopSingleOrderQuery) error {
    r.singleOrderQuery = singleOrderQuery
    r.Set("single_order_query", singleOrderQuery)
    return nil
}

func (r AliexpressTradeDsOrderGetRequest) GetSingleOrderQuery() *AeopSingleOrderQuery {
    return r.singleOrderQuery
}

