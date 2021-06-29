package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单预订接口 APIRequest
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

func NewAlitripMerchantGalaxyOrderBookRequest() *AlitripMerchantGalaxyOrderBookRequest{
    return &AlitripMerchantGalaxyOrderBookRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyOrderBookRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.book"
}

func (r AlitripMerchantGalaxyOrderBookRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyOrderBookRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyOrderBookRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyOrderBookRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyOrderBookRequest) GetToken() string {
    return r.token
}

func (r *AlitripMerchantGalaxyOrderBookRequest) SetOrderParam(orderParam *CreateOrderParam) error {
    r.orderParam = orderParam
    r.Set("order_param", orderParam)
    return nil
}

func (r AlitripMerchantGalaxyOrderBookRequest) GetOrderParam() *CreateOrderParam {
    return r.orderParam
}

func (r *AlitripMerchantGalaxyOrderBookRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r AlitripMerchantGalaxyOrderBookRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *AlitripMerchantGalaxyOrderBookRequest) SetSourceQuery(sourceQuery string) error {
    r.sourceQuery = sourceQuery
    r.Set("source_query", sourceQuery)
    return nil
}

func (r AlitripMerchantGalaxyOrderBookRequest) GetSourceQuery() string {
    return r.sourceQuery
}

