package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单列表查看(卖家视角) APIRequest
alibaba.lst.trade.seller.order.list.query

卖家视角订单查询，查询授权经销商订单列表
*/
type AlibabaLstTradeSellerOrderListQueryRequest struct {
    model.Params

    // 入参
    param   *LstTradeGetSellerOrderListParam 

}

func NewAlibabaLstTradeSellerOrderListQueryRequest() *AlibabaLstTradeSellerOrderListQueryRequest{
    return &AlibabaLstTradeSellerOrderListQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeSellerOrderListQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.order.list.query"
}

func (r AlibabaLstTradeSellerOrderListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeSellerOrderListQueryRequest) SetParam(param *LstTradeGetSellerOrderListParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaLstTradeSellerOrderListQueryRequest) GetParam() *LstTradeGetSellerOrderListParam {
    return r.param
}

