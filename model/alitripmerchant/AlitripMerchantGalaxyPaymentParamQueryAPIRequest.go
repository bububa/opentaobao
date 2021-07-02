package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyPaymentParamQueryAPIRequest 星河-支付参数查询接口 API请求
// alitrip.merchant.galaxy.payment.param.query
//
// 获取微信小程序的支付参数，供微信小程序调起支付请求。参数都为校验字段，不涉及用户隐私数据
type AlitripMerchantGalaxyPaymentParamQueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户校验token
	_token string
	// 订单编号
	_orderId string
}

// NewAlitripMerchantGalaxyPaymentParamQueryRequest 初始化AlitripMerchantGalaxyPaymentParamQueryAPIRequest对象
func NewAlitripMerchantGalaxyPaymentParamQueryRequest() *AlitripMerchantGalaxyPaymentParamQueryAPIRequest {
	return &AlitripMerchantGalaxyPaymentParamQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyPaymentParamQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.payment.param.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyPaymentParamQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyPaymentParamQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyPaymentParamQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户校验token
func (r *AlitripMerchantGalaxyPaymentParamQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyPaymentParamQueryAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *AlitripMerchantGalaxyPaymentParamQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyPaymentParamQueryAPIRequest) GetOrderId() string {
	return r._orderId
}
