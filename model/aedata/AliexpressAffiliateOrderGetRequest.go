package aedata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE流量订单详情获取API API请求
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

// 初始化AliexpressAffiliateOrderGetRequest对象
func NewAliexpressAffiliateOrderGetRequest() *AliexpressAffiliateOrderGetRequest{
    return &AliexpressAffiliateOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateOrderGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateOrderGetRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateOrderGetRequest) GetAppSignature() string {
    return r.appSignature
}
// Fields Setter
// 返回的字段列表
func (r *AliexpressAffiliateOrderGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateOrderGetRequest) GetFields() string {
    return r.fields
}
// OrderIds Setter
// 订单ID列表，以逗号分隔，当前只支持子订单ID查询
func (r *AliexpressAffiliateOrderGetRequest) SetOrderIds(orderIds string) error {
    r.orderIds = orderIds
    r.Set("order_ids", orderIds)
    return nil
}

// OrderIds Getter
func (r AliexpressAffiliateOrderGetRequest) GetOrderIds() string {
    return r.orderIds
}
