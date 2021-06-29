package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-取消预订 APIRequest
alitrip.merchant.galaxy.order.cancel

雅高酒店用户使用该接口，取消酒店预订
*/
type AlitripMerchantGalaxyOrderCancelRequest struct {
    model.Params

    // 租户标识
    tenantKey   string 

    // 用户登录标识
    token   string 

    // 订单编号
    orderId   string 

    // 退款原因
    reason   string 

}

func NewAlitripMerchantGalaxyOrderCancelRequest() *AlitripMerchantGalaxyOrderCancelRequest{
    return &AlitripMerchantGalaxyOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyOrderCancelRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.cancel"
}

func (r AlitripMerchantGalaxyOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyOrderCancelRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyOrderCancelRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyOrderCancelRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyOrderCancelRequest) GetToken() string {
    return r.token
}

func (r *AlitripMerchantGalaxyOrderCancelRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlitripMerchantGalaxyOrderCancelRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlitripMerchantGalaxyOrderCancelRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlitripMerchantGalaxyOrderCancelRequest) GetReason() string {
    return r.reason
}

