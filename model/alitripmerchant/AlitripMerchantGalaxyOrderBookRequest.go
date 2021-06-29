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
func (r *AlitripMerchantGalaxyOrderBookRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetTenantKey() string {
    return r._tenantKey
}
// Token Setter
// 用户登录token
func (r *AlitripMerchantGalaxyOrderBookRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetToken() string {
    return r._token
}
// OrderParam Setter
// 预订参数
func (r *AlitripMerchantGalaxyOrderBookRequest) SetOrderParam(_orderParam *CreateOrderParam) error {
    r._orderParam = _orderParam
    r.Set("order_param", _orderParam)
    return nil
}

// OrderParam Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetOrderParam() *CreateOrderParam {
    return r._orderParam
}
// OrderCode Setter
// 订单编号
func (r *AlitripMerchantGalaxyOrderBookRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetOrderCode() string {
    return r._orderCode
}
// SourceQuery Setter
// 广告追踪信息
func (r *AlitripMerchantGalaxyOrderBookRequest) SetSourceQuery(_sourceQuery string) error {
    r._sourceQuery = _sourceQuery
    r.Set("source_query", _sourceQuery)
    return nil
}

// SourceQuery Getter
func (r AlitripMerchantGalaxyOrderBookRequest) GetSourceQuery() string {
    return r._sourceQuery
}
