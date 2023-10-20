package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest v5.0付费会员卡开发订单开票信息查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query
//
// v5.0付费会员卡开发订单开票信息查询
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest struct {
	model.Params
	// 831232110dasdas
	_token string
	// 831232110
	_tenantKey string
	// 订单id
	_orderId string
}

// NewAlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryRequest 初始化AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryRequest() *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest {
	return &AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 831232110dasdas
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) GetToken() string {
	return r._token
}

// SetTenantKey is TenantKey Setter
// 831232110
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest) GetOrderId() string {
	return r._orderId
}
