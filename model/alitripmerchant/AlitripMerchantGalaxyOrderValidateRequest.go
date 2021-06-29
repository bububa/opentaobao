package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单试单接口 API请求
alitrip.merchant.galaxy.order.validate

根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息
*/
type AlitripMerchantGalaxyOrderValidateRequest struct {
    model.Params
    // 租户身份信息
    tenantKey   string
    // 试单参数
    validateOrderParam   *ValidateOrderParam
    // 用户标识
    token   string
}

// 初始化AlitripMerchantGalaxyOrderValidateRequest对象
func NewAlitripMerchantGalaxyOrderValidateRequest() *AlitripMerchantGalaxyOrderValidateRequest{
    return &AlitripMerchantGalaxyOrderValidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderValidateRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.validate"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOrderValidateRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderValidateRequest) GetTenantKey() string {
    return r.tenantKey
}
// ValidateOrderParam Setter
// 试单参数
func (r *AlitripMerchantGalaxyOrderValidateRequest) SetValidateOrderParam(validateOrderParam *ValidateOrderParam) error {
    r.validateOrderParam = validateOrderParam
    r.Set("validate_order_param", validateOrderParam)
    return nil
}

// ValidateOrderParam Getter
func (r AlitripMerchantGalaxyOrderValidateRequest) GetValidateOrderParam() *ValidateOrderParam {
    return r.validateOrderParam
}
// Token Setter
// 用户标识
func (r *AlitripMerchantGalaxyOrderValidateRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderValidateRequest) GetToken() string {
    return r.token
}
