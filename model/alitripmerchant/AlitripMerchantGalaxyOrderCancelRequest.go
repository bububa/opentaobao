package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-取消预订 API请求
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

// 初始化AlitripMerchantGalaxyOrderCancelRequest对象
func NewAlitripMerchantGalaxyOrderCancelRequest() *AlitripMerchantGalaxyOrderCancelRequest{
    return &AlitripMerchantGalaxyOrderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderCancelRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetTenantKey() string {
    return r.tenantKey
}
// Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetToken() string {
    return r.token
}
// OrderId Setter
// 订单编号
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetOrderId() string {
    return r.orderId
}
// Reason Setter
// 退款原因
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetReason() string {
    return r.reason
}
