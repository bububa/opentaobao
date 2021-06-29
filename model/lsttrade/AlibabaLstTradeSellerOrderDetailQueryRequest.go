package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情查看(卖家视角) APIRequest
alibaba.lst.trade.seller.order.detail.query

订单详情查看(卖家视角)
*/
type AlibabaLstTradeSellerOrderDetailQueryRequest struct {
    model.Params

    // 入参
    param   *LstTradeGetSellerOrderListParam 

}

func NewAlibabaLstTradeSellerOrderDetailQueryRequest() *AlibabaLstTradeSellerOrderDetailQueryRequest{
    return &AlibabaLstTradeSellerOrderDetailQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeSellerOrderDetailQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.order.detail.query"
}

func (r AlibabaLstTradeSellerOrderDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeSellerOrderDetailQueryRequest) SetParam(param *LstTradeGetSellerOrderListParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaLstTradeSellerOrderDetailQueryRequest) GetParam() *LstTradeGetSellerOrderListParam {
    return r.param
}

