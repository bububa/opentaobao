package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-支付参数查询接口 API请求
alitrip.merchant.galaxy.payment.param.query

获取微信小程序的支付参数，供微信小程序调起支付请求。参数都为校验字段，不涉及用户隐私数据
*/
type AlitripMerchantGalaxyPaymentParamQueryRequest struct {
    model.Params
    // 租户身份信息
    tenantKey   string
    // 用户校验token
    token   string
    // 订单编号
    orderId   string
}

// 初始化AlitripMerchantGalaxyPaymentParamQueryRequest对象
func NewAlitripMerchantGalaxyPaymentParamQueryRequest() *AlitripMerchantGalaxyPaymentParamQueryRequest{
    return &AlitripMerchantGalaxyPaymentParamQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.payment.param.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetTenantKey() string {
    return r.tenantKey
}
// Token Setter
// 用户校验token
func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetToken() string {
    return r.token
}
// OrderId Setter
// 订单编号
func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetOrderId() string {
    return r.orderId
}
