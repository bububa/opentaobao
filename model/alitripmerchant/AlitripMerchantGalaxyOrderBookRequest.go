package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单预订接口 API请求
alitrip.merchant.galaxy.order.book

为星河酒店解决方案的C端用户提供酒店预订能力
*/
type AlitripMerchantGalaxyOrderBookRequest struct {
    model.Params
    // 租户身份信息
    tenantKey   string
    // 用户登录token
    token   string
    // 预订参数
    orderParam   *CreateOrderParam
    // 订单编号
    orderCode   string
    // 广告追踪信息
    sourceQuery   string
}

// 初始化AlitripMerchantGalaxyOrderBookRequest对象
func NewAlitripMerchantGalaxyOrderBookRequest() *AlitripMerchantGalaxyOrderBookRequest{
    return &AlitripMerchantGalaxyOrderBookRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderBookRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.book"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderBookRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOrderBookRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetTenantKey() string {
    return r.tenantKey
}
// Token Setter
// 用户登录token
func (r *AlitripMerchantGalaxyOrderBookRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetToken() string {
    return r.token
}
// OrderParam Setter
// 预订参数
func (r *AlitripMerchantGalaxyOrderBookRequest) SetOrderParam(orderParam *CreateOrderParam) error {
    r.orderParam = orderParam
    r.Set("order_param", orderParam)
    return nil
}

// OrderParam Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetOrderParam() *CreateOrderParam {
    return r.orderParam
}
// OrderCode Setter
// 订单编号
func (r *AlitripMerchantGalaxyOrderBookRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetOrderCode() string {
    return r.orderCode
}
// SourceQuery Setter
// 广告追踪信息
func (r *AlitripMerchantGalaxyOrderBookRequest) SetSourceQuery(sourceQuery string) error {
    r.sourceQuery = sourceQuery
    r.Set("source_query", sourceQuery)
    return nil
}

// SourceQuery Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetSourceQuery() string {
    return r.sourceQuery
}
