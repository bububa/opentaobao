package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest 会员权益卡身份识别二维码图片 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode
//
// 会员权益卡身份识别二维码图片
type AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest struct {
	model.Params
	// 123
	_tenantKey string
	// 100
	_token string
	// 卡号
	_cardNumber string
	// 权益卡号
	_memberVoucherCardID string
}

// NewAlitripmerchantgalaxyderbymembervouchercardshowqrcodeRequest 初始化AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardshowqrcodeRequest() *AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 123
func (r *AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 100
func (r *AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) GetToken() string {
	return r._token
}

// SetCardNumber is CardNumber Setter
// 卡号
func (r *AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) SetCardNumber(_cardNumber string) error {
	r._cardNumber = _cardNumber
	r.Set("card_number", _cardNumber)
	return nil
}

// GetCardNumber CardNumber Getter
func (r AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) GetCardNumber() string {
	return r._cardNumber
}

// SetMemberVoucherCardID is MemberVoucherCardID Setter
// 权益卡号
func (r *AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) SetMemberVoucherCardID(_memberVoucherCardID string) error {
	r._memberVoucherCardID = _memberVoucherCardID
	r.Set("member_voucher_card_i_d", _memberVoucherCardID)
	return nil
}

// GetMemberVoucherCardID MemberVoucherCardID Getter
func (r AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest) GetMemberVoucherCardID() string {
	return r._memberVoucherCardID
}
