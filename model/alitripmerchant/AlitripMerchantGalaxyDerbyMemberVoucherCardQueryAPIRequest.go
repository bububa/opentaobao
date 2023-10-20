package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest 德比付费会员卡查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.query
//
// 德比付费会员卡查询
type AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 1
	_token string
	// 1
	_memberVoucherCardID string
}

// NewAlitripmerchantgalaxyderbymembervouchercardqueryRequest 初始化AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardqueryRequest() *AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) GetToken() string {
	return r._token
}

// SetMemberVoucherCardID is MemberVoucherCardID Setter
// 1
func (r *AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) SetMemberVoucherCardID(_memberVoucherCardID string) error {
	r._memberVoucherCardID = _memberVoucherCardID
	r.Set("member_voucher_card_i_d", _memberVoucherCardID)
	return nil
}

// GetMemberVoucherCardID MemberVoucherCardID Getter
func (r AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest) GetMemberVoucherCardID() string {
	return r._memberVoucherCardID
}
