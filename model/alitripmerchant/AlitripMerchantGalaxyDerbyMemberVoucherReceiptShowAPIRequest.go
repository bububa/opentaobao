package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest v5.0付费会员卡开发发票图片展示 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.show
//
// v5.0付费会员卡开发发票图片展示
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 订单号
	_orderId string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowRequest() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.show"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest) GetOrderId() string {
	return r._orderId
}
