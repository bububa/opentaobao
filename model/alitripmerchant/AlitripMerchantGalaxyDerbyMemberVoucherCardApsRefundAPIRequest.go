package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest Aps退券通知接口 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
type AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 请求String
	_derbyBody string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetDerbyBody is DerbyBody Setter
// 请求String
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest) SetDerbyBody(_derbyBody string) error {
	r._derbyBody = _derbyBody
	r.Set("derby_body", _derbyBody)
	return nil
}

// GetDerbyBody DerbyBody Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest) GetDerbyBody() string {
	return r._derbyBody
}
