package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest 权益卡订单激活 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.active
//
// 权益卡订单激活
type AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 会员卡号
	_cardNumber string
	// 权益卡号
	_memberVoucherCardID string
}

// NewAlitripmerchantgalaxyderbymembervouchercardactiveRequest 初始化AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardactiveRequest() *AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.active"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) GetToken() string {
	return r._token
}

// SetCardNumber is CardNumber Setter
// 会员卡号
func (r *AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) SetCardNumber(_cardNumber string) error {
	r._cardNumber = _cardNumber
	r.Set("card_number", _cardNumber)
	return nil
}

// GetCardNumber CardNumber Getter
func (r AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) GetCardNumber() string {
	return r._cardNumber
}

// SetMemberVoucherCardID is MemberVoucherCardID Setter
// 权益卡号
func (r *AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) SetMemberVoucherCardID(_memberVoucherCardID string) error {
	r._memberVoucherCardID = _memberVoucherCardID
	r.Set("member_voucher_card_i_d", _memberVoucherCardID)
	return nil
}

// GetMemberVoucherCardID MemberVoucherCardID Getter
func (r AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest) GetMemberVoucherCardID() string {
	return r._memberVoucherCardID
}
