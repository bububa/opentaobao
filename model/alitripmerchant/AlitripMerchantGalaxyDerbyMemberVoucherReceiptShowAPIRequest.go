package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest v5.0付费会员卡开发发票图片展示 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.show
//
// v5.0付费会员卡开发发票图片展示
type AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 订单号
	_orderId string
}

// NewAlitripmerchantgalaxyderbymembervoucherreceiptshowRequest 初始化AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervoucherreceiptshowRequest() *AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest {
	return &AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.show"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptshowAPIRequest) GetOrderId() string {
	return r._orderId
}
