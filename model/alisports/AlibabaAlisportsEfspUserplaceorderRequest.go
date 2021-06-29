package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户完成支付同步订单 APIRequest
alibaba.alisports.efsp.userplaceorder

用户完成支付同步订单
*/
type AlibabaAlisportsEfspUserplaceorderRequest struct {
    model.Params

    // 青橙订单的json
    orderJson   string 

}

func NewAlibabaAlisportsEfspUserplaceorderRequest() *AlibabaAlisportsEfspUserplaceorderRequest{
    return &AlibabaAlisportsEfspUserplaceorderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsEfspUserplaceorderRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.userplaceorder"
}

func (r AlibabaAlisportsEfspUserplaceorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsEfspUserplaceorderRequest) SetOrderJson(orderJson string) error {
    r.orderJson = orderJson
    r.Set("order_json", orderJson)
    return nil
}

func (r AlibabaAlisportsEfspUserplaceorderRequest) GetOrderJson() string {
    return r.orderJson
}

