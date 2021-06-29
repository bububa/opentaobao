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
    _appSignature   string
    // 返回的字段列表
    _fields   string
    // 订单ID列表，以逗号分隔，当前只支持子订单ID查询
    _orderIds   string
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
func (r *AliexpressAffiliateOrderGetRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateOrderGetRequest) GetAppSignature() string {
    return r._appSignature
}
// Fields Setter
// 返回的字段列表
func (r *AliexpressAffiliateOrderGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateOrderGetRequest) GetFields() string {
    return r._fields
}
// OrderIds Setter
// 订单ID列表，以逗号分隔，当前只支持子订单ID查询
func (r *AliexpressAffiliateOrderGetRequest) SetOrderIds(_orderIds string) error {
    r._orderIds = _orderIds
    r.Set("order_ids", _orderIds)
    return nil
}

// OrderIds Getter
func (r AliexpressAffiliateOrderGetRequest) GetOrderIds() string {
    return r._orderIds
}
