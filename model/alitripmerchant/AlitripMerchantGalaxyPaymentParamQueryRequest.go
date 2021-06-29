package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-支付参数查询接口 APIRequest
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

func NewAlitripMerchantGalaxyPaymentParamQueryRequest() *AlitripMerchantGalaxyPaymentParamQueryRequest{
    return &AlitripMerchantGalaxyPaymentParamQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.payment.param.query"
}

func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetToken() string {
    return r.token
}

func (r *AlitripMerchantGalaxyPaymentParamQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlitripMerchantGalaxyPaymentParamQueryRequest) GetOrderId() string {
    return r.orderId
}

