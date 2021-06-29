package aedata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE流量订单详情获取API APIRequest
aliexpress.affiliate.order.get

联盟推广订单效果获取API
*/
type AliexpressAffiliateOrderGetRequest struct {
    model.Params

    // 安全签名
    appSignature   string 

    // 返回的字段列表
    fields   string 

    // 订单ID列表，以逗号分隔，当前只支持子订单ID查询
    orderIds   string 

}

func NewAliexpressAffiliateOrderGetRequest() *AliexpressAffiliateOrderGetRequest{
    return &AliexpressAffiliateOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateOrderGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.order.get"
}

func (r AliexpressAffiliateOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateOrderGetRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateOrderGetRequest) GetAppSignature() string {
    return r.appSignature
}

func (r *AliexpressAffiliateOrderGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateOrderGetRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateOrderGetRequest) SetOrderIds(orderIds string) error {
    r.orderIds = orderIds
    r.Set("order_ids", orderIds)
    return nil
}

func (r AliexpressAffiliateOrderGetRequest) GetOrderIds() string {
    return r.orderIds
}

