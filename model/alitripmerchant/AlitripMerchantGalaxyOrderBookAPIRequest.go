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
type AlitripMerchantGalaxyOrderBookAPIRequest struct {
    model.Params
    // 租户身份信息
    _tenantKey   string
    // 用户登录token
    _token   string
    // 预订参数
    _orderParam   *CreateOrderParam
    // 订单编号
    _orderCode   string
    // 广告追踪信息
    _sourceQuery   string
}

// 初始化AlitripMerchantGalaxyOrderBookAPIRequest对象
func NewAlitripMerchantGalaxyOrderBookRequest() *AlitripMerchantGalaxyOrderBookAPIRequest{
    return &AlitripMerchantGalaxyOrderBookAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderBookAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.book"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderBookAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOrderBookAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderBookAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
// Token Setter
// 用户登录token
func (r *AlitripMerchantGalaxyOrderBookAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderBookAPIRequest) GetToken() string {
    return r._token
}
// OrderParam Setter
// 预订参数
func (r *AlitripMerchantGalaxyOrderBookAPIRequest) SetOrderParam(_orderParam *CreateOrderParam) error {
    r._orderParam = _orderParam
    r.Set("order_param", _orderParam)
    return nil
}

// OrderParam Getter
func (r AlitripMerchantGalaxyOrderBookAPIRequest) GetOrderParam() *CreateOrderParam {
    return r._orderParam
}
// OrderCode Setter
// 订单编号
func (r *AlitripMerchantGalaxyOrderBookAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r AlitripMerchantGalaxyOrderBookAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// SourceQuery Setter
// 广告追踪信息
func (r *AlitripMerchantGalaxyOrderBookAPIRequest) SetSourceQuery(_sourceQuery string) error {
    r._sourceQuery = _sourceQuery
    r.Set("source_query", _sourceQuery)
    return nil
}

// SourceQuery Getter
func (r AlitripMerchantGalaxyOrderBookAPIRequest) GetSourceQuery() string {
    return r._sourceQuery
}
