package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest 订单详情查询接口 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query
//
// 订单详情查询接口
type AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 订单号
	_orderId string
}

// NewAlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryRequest 初始化AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryRequest() *AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest) GetOrderId() string {
	return r._orderId
}
