package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest 德比付费会员卡查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.query
//
// 德比付费会员卡查询
type AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 1
	_token string
	// 1
	_memberVoucherCardID string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardQueryRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardQueryRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) GetToken() string {
	return r._token
}

// SetMemberVoucherCardID is MemberVoucherCardID Setter
// 1
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) SetMemberVoucherCardID(_memberVoucherCardID string) error {
	r._memberVoucherCardID = _memberVoucherCardID
	r.Set("member_voucher_card_i_d", _memberVoucherCardID)
	return nil
}

// GetMemberVoucherCardID MemberVoucherCardID Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest) GetMemberVoucherCardID() string {
	return r._memberVoucherCardID
}
