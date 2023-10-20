package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxypaymentparamqueryAPIRequest 星河-支付参数查询接口 API请求
// alitrip.merchant.galaxy.payment.param.query
//
// 获取微信小程序的支付参数，供微信小程序调起支付请求。参数都为校验字段，不涉及用户隐私数据
type AlitripmerchantgalaxypaymentparamqueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户校验token
	_token string
	// 订单编号
	_orderId string
}

// NewAlitripmerchantgalaxypaymentparamqueryRequest 初始化AlitripmerchantgalaxypaymentparamqueryAPIRequest对象
func NewAlitripmerchantgalaxypaymentparamqueryRequest() *AlitripmerchantgalaxypaymentparamqueryAPIRequest {
	return &AlitripmerchantgalaxypaymentparamqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxypaymentparamqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.payment.param.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxypaymentparamqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxypaymentparamqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxypaymentparamqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxypaymentparamqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户校验token
func (r *AlitripmerchantgalaxypaymentparamqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxypaymentparamqueryAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *AlitripmerchantgalaxypaymentparamqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripmerchantgalaxypaymentparamqueryAPIRequest) GetOrderId() string {
	return r._orderId
}
