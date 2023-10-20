package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest v5.0德比付费会员卡通知 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.change.callback
//
// v5.0德比付费会员卡通知
type AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 回调方法
	_derbyMethod string
	// 1
	_derbyBody string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.change.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetDerbyMethod is DerbyMethod Setter
// 回调方法
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) SetDerbyMethod(_derbyMethod string) error {
	r._derbyMethod = _derbyMethod
	r.Set("derby_method", _derbyMethod)
	return nil
}

// GetDerbyMethod DerbyMethod Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) GetDerbyMethod() string {
	return r._derbyMethod
}

// SetDerbyBody is DerbyBody Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) SetDerbyBody(_derbyBody string) error {
	r._derbyBody = _derbyBody
	r.Set("derby_body", _derbyBody)
	return nil
}

// GetDerbyBody DerbyBody Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest) GetDerbyBody() string {
	return r._derbyBody
}
