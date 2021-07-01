package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-单个订单详细信息查询 API请求
alitrip.merchant.galaxy.order.query

为用户提供酒店订单的详细信息查询能力
*/
type AlitripMerchantGalaxyOrderQueryAPIRequest struct {
    model.Params
    // 租户标识
    _tenantKey   string
    // 用户登录标识
    _token   string
    // 订单号
    _orderId   string
}

// 初始化AlitripMerchantGalaxyOrderQueryAPIRequest对象
func NewAlitripMerchantGalaxyOrderQueryRequest() *AlitripMerchantGalaxyOrderQueryAPIRequest{
    return &AlitripMerchantGalaxyOrderQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderQueryAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
// Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyOrderQueryAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetToken() string {
    return r._token
}
// OrderId Setter
// 订单号
func (r *AlitripMerchantGalaxyOrderQueryAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetOrderId() string {
    return r._orderId
}
