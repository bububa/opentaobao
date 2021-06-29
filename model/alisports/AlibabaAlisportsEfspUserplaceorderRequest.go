package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户完成支付同步订单 API请求
alibaba.alisports.efsp.userplaceorder

用户完成支付同步订单
*/
type AlibabaAlisportsEfspUserplaceorderRequest struct {
    model.Params
    // 青橙订单的json
    orderJson   string
}

// 初始化AlibabaAlisportsEfspUserplaceorderRequest对象
func NewAlibabaAlisportsEfspUserplaceorderRequest() *AlibabaAlisportsEfspUserplaceorderRequest{
    return &AlibabaAlisportsEfspUserplaceorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspUserplaceorderRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.userplaceorder"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspUserplaceorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderJson Setter
// 青橙订单的json
func (r *AlibabaAlisportsEfspUserplaceorderRequest) SetOrderJson(orderJson string) error {
    r.orderJson = orderJson
    r.Set("order_json", orderJson)
    return nil
}

// OrderJson Getter
func (r AlibabaAlisportsEfspUserplaceorderRequest) GetOrderJson() string {
    return r.orderJson
}
