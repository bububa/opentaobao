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
    _tenantKey   string
    // 用户校验token
    _token   string
    // 订单编号
    _orderId   string
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
func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetTenantKey() string {
    return r._tenantKey
}
// Token Setter
// 用户校验token
func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetToken() string {
    return r._token
}
// OrderId Setter
// 订单编号
func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetOrderId() string {
    return r._orderId
}
