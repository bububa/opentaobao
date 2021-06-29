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
    _tenantKey   string
    // 用户登录标识
    _token   string
    // 订单编号
    _orderId   string
    // 退款原因
    _reason   string
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
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetTenantKey() string {
    return r._tenantKey
}
// Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetToken() string {
    return r._token
}
// OrderId Setter
// 订单编号
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetOrderId() string {
    return r._orderId
}
// Reason Setter
// 退款原因
func (r *AlitripMerchantGalaxyOrderCancelRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlitripMerchantGalaxyOrderCancelRequest) GetReason() string {
    return r._reason
}
