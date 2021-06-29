package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-单个订单详细信息查询 APIRequest
alitrip.merchant.galaxy.order.query

为用户提供酒店订单的详细信息查询能力
*/
type AlitripMerchantGalaxyOrderQueryRequest struct {
    model.Params

    // 租户标识
    tenantKey   string 

    // 用户登录标识
    token   string 

    // 订单号
    orderId   string 

}

func NewAlitripMerchantGalaxyOrderQueryRequest() *AlitripMerchantGalaxyOrderQueryRequest{
    return &AlitripMerchantGalaxyOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyOrderQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.query"
}

func (r AlitripMerchantGalaxyOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyOrderQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyOrderQueryRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyOrderQueryRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyOrderQueryRequest) GetToken() string {
    return r.token
}

func (r *AlitripMerchantGalaxyOrderQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlitripMerchantGalaxyOrderQueryRequest) GetOrderId() string {
    return r.orderId
}

